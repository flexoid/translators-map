package services

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/flexoid/translators-map-go/ent"
	"github.com/flexoid/translators-map-go/ent/translator"
	"github.com/flexoid/translators-map-go/internal/maps"
	"github.com/flexoid/translators-map-go/internal/scraper"
	"go.uber.org/zap"
)

type Scraper struct {
	db                *ent.Client
	logger            *zap.SugaredLogger
	geocoding         *maps.Geocoding
	mapsBackendAPIKey string
}

func NewScraper(db *ent.Client, logger *zap.SugaredLogger, mapsBackendAPIKey string) *Scraper {
	return &Scraper{db: db, logger: logger, mapsBackendAPIKey: mapsBackendAPIKey}
}

func (s *Scraper) Run() {
	var err error
	s.geocoding, err = maps.NewGeocoding(s.mapsBackendAPIKey)
	if err != nil {
		s.logger.Fatalf("failed to setup geocoding client: %v", err)
	}

	languages, err := scraper.ScrapeLanguages(s.logger)
	if err != nil {
		s.logger.Fatalf("Failed to scrape languages: %v", err)
		return
	}

	for _, language := range languages {
		err := scraper.ScrapeTranslators(s.logger, language, func(t scraper.Translator) {
			_, err := s.handleTranslator(t)
			if err != nil {
				s.logger.Errorf("Failed to save translator to db: %v", err)
			}
		})
		if err != nil {
			s.logger.Fatalf("Error while scraping translators: %v", err)
		}
	}
}

func (s *Scraper) handleTranslator(trans scraper.Translator) (*ent.Translator, error) {
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

	err := s.fillLocation(context.TODO(), creator.Mutation(), trans.Address)
	if err != nil {
		return nil, err
	}

	model, err := creator.Save(context.TODO())

	if err != nil {
		return nil, fmt.Errorf("failed to create translator record: %w", err)
	}

	s.logger.Debugw("Created translator record", "model", model.String())

	return model, nil
}

func (s *Scraper) updateTranslator(model *ent.Translator, trans scraper.Translator) (*ent.Translator, error) {
	addressSum := s.hashSumFromString(trans.Address)
	s.logger.Debugf("Comparing address hashsum from database %x to scraped one %x",
		addressSum, model.AddressSha)

	if bytes.Equal(model.AddressSha, addressSum) {
		s.logger.Debugf("Address didn't change, skipping update")
		return model, nil
	}

	updater := model.Update()

	err := s.fillLocation(context.TODO(), updater.Mutation(), trans.Address)
	if err != nil {
		return nil, err
	}

	model, err = updater.Save(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to update translator record: %w", err)
	}

	s.logger.Debugw("Updated translator record", "model", model.String())

	return model, nil
}

func (s *Scraper) hashSumFromString(str string) []byte {
	sum := sha256.Sum256([]byte(str))
	return sum[:]
}

func (s *Scraper) fillLocation(ctx context.Context, m *ent.TranslatorMutation, address string) error {
	geocodingResult, err := s.geocoding.GetCoordinatesForAddress(ctx, address)
	if err != nil {
		return fmt.Errorf("geocoding error: %v", err)
	}

	lat := geocodingResult.Geometry.Location.Lat
	lng := geocodingResult.Geometry.Location.Lng
	addressSum := s.hashSumFromString(address)
	s.logger.Debugw("Got location for address", "address", address,
		"address_sha", hex.EncodeToString(addressSum), "latitude", lat, "longitude", lng)

	m.SetAddress(address)
	m.SetAddressSha(addressSum)
	m.SetLatitude(lat)
	m.SetLongitude(lng)

	return nil
}
