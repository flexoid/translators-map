package services

import (
	"context"
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
		Where(translator.Name(trans.Name), translator.Language(trans.Language.Name)).
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

	creator.SetName(trans.Name)
	creator.SetLanguage(trans.Language.Name)
	s.setModelAttrs(creator.Mutation(), trans)

	geocodingResult, err := s.geocoding.GetCoordinatesForAddress(context.TODO(), trans.Address)
	if err != nil {
		return nil, fmt.Errorf("geocoding error: %v", err)
	}

	lat := geocodingResult.Geometry.Location.Lat
	lng := geocodingResult.Geometry.Location.Lng
	s.logger.Debugw("Got location for address", "address", trans.Address, "latitude", lat, "longitude", lng)

	creator.
		SetLatitude(lat).
		SetLongitude(lng)

	model, err := creator.Save(context.TODO())

	if err != nil {
		return nil, fmt.Errorf("failed to create translator record: %w", err)
	}

	s.logger.Debugw("Created translator record",
		"id", model.ID, "name", model.Name,
		"language", model.Language, "model", model.String())

	return model, nil
}

func (s *Scraper) updateTranslator(model *ent.Translator, trans scraper.Translator) (*ent.Translator, error) {
	updater := model.Update()
	s.setModelAttrs(updater.Mutation(), trans)

	model, err := updater.Save(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to update translator record: %w", err)
	}

	s.logger.Debugw("Updated translator record",
		"id", model.ID, "name", model.Name,
		"language", model.Language, "model", model.String())

	return model, nil
}

func (s *Scraper) setModelAttrs(m *ent.TranslatorMutation, trans scraper.Translator) {
	m.SetAddress(trans.Address)

	m.SetContacts(trans.Contacts)
	m.SetDetailsURL(trans.DetailsURL)
}
