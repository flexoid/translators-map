package services

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/flexoid/translators-map-go/ent"
	"github.com/flexoid/translators-map-go/ent/translator"
	"github.com/flexoid/translators-map-go/internal/config"
	"github.com/flexoid/translators-map-go/internal/maps"
	"github.com/flexoid/translators-map-go/internal/metrics"
	"github.com/flexoid/translators-map-go/internal/scraper"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"go.uber.org/zap"
)

type Scraper struct {
	db        *ent.Client
	logger    *zap.SugaredLogger
	geocoding Geocoder
	metrics   *metrics.ScraperMetrics
}

type Geocoder interface {
	GeocodingForAddress(ctx context.Context, address string) (*maps.Result, error)
}

const ScraperRunInterval = 24 * time.Hour
const DeleteOldTranslatorsInterval = 3 * ScraperRunInterval

// Sets the threshold to consider scraper run successful.
const percentToSuccess = 0.9

func NewScraper(
	db *ent.Client,
	logger *zap.SugaredLogger,
	geocoding Geocoder,
	metricset *metrics.ScraperMetrics,
) *Scraper {
	return &Scraper{
		db:        db,
		logger:    logger,
		geocoding: geocoding,
		metrics:   metricset,
	}
}

func (s *Scraper) Run() {
	startTime := time.Now()
	successful := false
	translatorsScraped := 0
	translatorsFailed := 0

	s.metrics.ResetBeforeRun()
	defer func() {
		s.sendMetrics(startTime, successful)
	}()

	languages, err := scraper.ScrapeLanguages(s.logger)
	if err != nil {
		s.logger.Errorf("Failed to scrape languages: %v", err)
		return
	}

	s.logger.Debugf("Scraped %d languages", len(languages))
	s.metrics.LanguagesScraped.Set(float64(len(languages)))

	for _, language := range languages {
		err := scraper.ScrapeTranslators(s.logger, language, func(t scraper.Translator) {
			translatorsScraped++
			_, err := s.handleTranslator(t)
			if err != nil {
				s.logger.Errorf("Failed to save translator to db: %v", err)
				s.metrics.TranslatorsFailed.Inc()
				translatorsFailed++
			}
		})
		if err != nil {
			s.logger.Errorf("Error while scraping translators: %v", err)
			return
		}
	}

	successful = s.isRunSuccessful(translatorsScraped, translatorsFailed)
	s.metrics.SuccessTime.SetToCurrentTime()

	if successful {
		s.deleteOldTranslators()
	}
}

func (s *Scraper) handleTranslator(trans scraper.Translator) (*ent.Translator, error) {
	s.metrics.TranslatorsScraped.Inc()

	var model *ent.Translator
	var err error

	model, err = s.db.Translator.Query().
		Where(
			translator.ExternalID(trans.ID),
			translator.Language(trans.Language.Name)).
		Only(context.TODO())

	if ent.IsNotFound(err) {
		model, err = s.createTranslator(trans)
		if err != nil {
			return nil, err
		}
		return model, nil
	} else if ent.IsNotSingular(err) {
		return nil, fmt.Errorf("there is more than one translator: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("failed to query translator: %w", err)
	}

	s.updateTranslator(model, trans)
	return model, nil
}

func (s *Scraper) createTranslator(trans scraper.Translator) (*ent.Translator, error) {
	creator := s.db.Translator.Create()

	creator.SetExternalID(trans.ID)
	creator.SetLanguage(trans.Language.Name)
	creator.SetDetailsURL(trans.DetailsURL)
	s.fillInfo(context.TODO(), nil, creator.Mutation(), &trans)

	err := s.fillLocation(context.TODO(), creator.Mutation(), trans.Address)
	if err != nil {
		return nil, err
	}

	model, err := creator.Save(context.TODO())

	if err != nil {
		return nil, fmt.Errorf("failed to create translator record: %w", err)
	}

	s.logger.Debugw("Created translator record", "model", model.String())
	s.metrics.TranslatorsCreated.Inc()

	return model, nil
}

func (s *Scraper) updateTranslator(model *ent.Translator, trans scraper.Translator) (*ent.Translator, error) {
	updater := model.Update()
	s.fillInfo(context.TODO(), model, updater.Mutation(), &trans)

	addressSum := s.hashSumFromString(trans.Address)
	s.logger.Debugf("Comparing address hashsum from database %x to scraped one %x",
		addressSum, model.AddressSha)

	if !bytes.Equal(model.AddressSha, addressSum) {
		s.logger.Debug("Address changed, updating location with geocoding API")

		err := s.fillLocation(context.TODO(), updater.Mutation(), trans.Address)
		if err != nil {
			return nil, err
		}
	} else {
		s.logger.Debugf("Address didn't change, skipping update")
	}

	model, err := updater.Save(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to update translator record: %w", err)
	}

	s.logger.Debugw("Updated translator record", "model", model.String())
	s.metrics.TranslatorsUpdated.Inc()

	return model, nil
}

func (s *Scraper) hashSumFromString(str string) []byte {
	sum := sha256.Sum256([]byte(str))
	return sum[:]
}

func (s *Scraper) fillInfo(ctx context.Context, model *ent.Translator, m *ent.TranslatorMutation, translator *scraper.Translator) {
	if model == nil || model.Name != translator.Name {
		m.SetName(translator.Name)
	}
}

func (s *Scraper) fillLocation(ctx context.Context, m *ent.TranslatorMutation, address string) error {
	geocodingResult, err := s.geocoding.GeocodingForAddress(ctx, address)
	if err != nil {
		return fmt.Errorf("geocoding error: %v", err)
	}

	addressSum := s.hashSumFromString(address)
	s.logger.Debugw("Got geocoding result for address", "address", address,
		"address_sha", hex.EncodeToString(addressSum), "geocoding", geocodingResult)

	m.SetAddress(address)
	m.SetAddressSha(addressSum)
	m.SetLatitude(geocodingResult.Lat)
	m.SetLongitude(geocodingResult.Lng)
	m.SetCity(geocodingResult.City)
	m.SetAdministrativeArea(geocodingResult.AdministrativeArea)
	m.SetCountry(geocodingResult.Country)

	return nil
}

func (s *Scraper) isRunSuccessful(total, failed int) bool {
	if total == 0 {
		return false
	}

	successful := float64(total-failed) / float64(total)
	return successful >= percentToSuccess
}

// Delete translators that were not updated during last N runs.
func (s *Scraper) deleteOldTranslators() error {
	affected, err := s.db.Translator.Delete().Where(
		translator.UpdatedAtLT(time.Now().Add(-DeleteOldTranslatorsInterval)),
	).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete old translators: %w", err)
	}

	s.logger.Debugf("Deleted %d translators", affected)
	s.metrics.TranslatorsDeleted.Set(float64(affected))
	return nil
}

func (s *Scraper) sendMetrics(startTime time.Time, successful bool) {
	s.metrics.CompletionTime.SetToCurrentTime()
	s.metrics.Duration.Set(time.Since(startTime).Seconds())

	if config.CLI.MetricsPushgatewayURL == "" || config.CLI.MetricsInstance == "" {
		s.logger.Debug("Skipping sending metrics to pushgateway")
		return
	}

	pusher := push.New(config.CLI.MetricsPushgatewayURL, "translators_map").
		Gatherer(prometheus.DefaultGatherer).
		Grouping("instance", config.CLI.MetricsInstance)

	if successful {
		// Include success timestamp collector only if scraper succeeded.
		pusher = pusher.Collector(s.metrics.SuccessTime)
	}

	err := pusher.Add()
	if err != nil {
		s.logger.Errorf("Failed to push metrics to pushgateway: %v", err)
	}
}
