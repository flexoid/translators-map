package main

import (
	"github.com/flexoid/translators-map-go/internal/scraper"
	"go.uber.org/zap"
)

func main() {
	logCfg := zap.NewDevelopmentConfig()
	logCfg.Level.SetLevel(zap.DebugLevel)

	zapLogger := zap.Must(logCfg.Build())
	defer zapLogger.Sync() // flushes buffer, if any
	logger := zapLogger.Sugar()

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
	}
}
