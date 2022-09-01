package main

import (
	"context"
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/flexoid/translators-map-go/ent"
	"github.com/flexoid/translators-map-go/ent/translator"
	"github.com/flexoid/translators-map-go/internal/scraper"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var CLI struct {
	DatabaseURL string `required:"" env:"DATABASE_URL" help:"Postgres database URL."`
}

func main() {
	_ = kong.Parse(&CLI)

	logCfg := zap.NewDevelopmentConfig()
	logCfg.Level.SetLevel(zap.DebugLevel)

	zapLogger := zap.Must(logCfg.Build())
	defer zapLogger.Sync() // flushes buffer, if any
	logger := zapLogger.Sugar()

	entClient, err := setupDatabase()
	if err != nil {
		logger.Fatalf("failed to setup database: %v", err)
	}

	languages, err := scraper.ScrapeLanguages(logger)
	if err != nil {
		logger.Fatalf("Failed to scrape languages: %v", err)
		return
	}

	// TODO: Scrape not only first 2 languages.
	for _, language := range languages[:2] {
		translators, err := scraper.ScrapeTranslators(logger, language)
		if err != nil {
			logger.Fatalf("Failed to scrape translators: %v", err)
		}

		logger.Debugf("Scraped %d translators", len(translators))

		for _, translator := range translators {
			_, err := handleTranslator(logger, entClient, translator)
			if err != nil {
				logger.Errorf("Failed to save translator to db: %v", err)
				continue
			}
		}
	}
}

func setupDatabase() (*ent.Client, error) {
	client, err := ent.Open("postgres", CLI.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	return client, nil
}

func handleTranslator(logger *zap.SugaredLogger, entClient *ent.Client, trans scraper.Translator) (*ent.Translator, error) {
	var model *ent.Translator
	var err error

	model, err = entClient.Translator.Query().
		Where(translator.Name(trans.Name), translator.Language(trans.Language.Name)).
		Only(context.TODO())

	if ent.IsNotFound(err) {
		model, err = createTranslator(logger, entClient, trans)
		if err != nil {
			return nil, err
		}
		return model, nil
	} else if ent.IsNotSingular(err) {
		return nil, fmt.Errorf("there is more than one translator: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("failed to query translator: %w", err)
	}

	updateTranslator(logger, model, trans)
	return model, nil
}

func createTranslator(logger *zap.SugaredLogger, entClient *ent.Client, trans scraper.Translator) (*ent.Translator, error) {
	creator := entClient.Translator.Create()

	creator.SetName(trans.Name)
	creator.SetLanguage(trans.Language.Name)
	setModelAttrs(creator.Mutation(), trans)

	model, err := creator.Save(context.TODO())

	if err != nil {
		return nil, fmt.Errorf("failed to create translator record: %w", err)
	}

	logger.Debugw("Created translator record",
		"id", model.ID, "name", model.Name,
		"language", model.Language, "model", model.String())

	return model, nil
}

func updateTranslator(logger *zap.SugaredLogger, model *ent.Translator, trans scraper.Translator) (*ent.Translator, error) {
	updater := model.Update()
	setModelAttrs(updater.Mutation(), trans)

	model, err := updater.Save(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to update translator record: %w", err)
	}

	logger.Debugw("Updated translator record",
		"id", model.ID, "name", model.Name,
		"language", model.Language, "model", model.String())

	return model, nil
}

func setModelAttrs(m *ent.TranslatorMutation, trans scraper.Translator) {
	m.SetAddress(trans.Address)
	m.SetContacts(trans.Contacts)
	m.SetDetailsURL(trans.DetailsURL)
}
