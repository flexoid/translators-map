package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"
	"github.com/flexoid/translators-map-go/ent"
	"github.com/flexoid/translators-map-go/internal/api"
	"github.com/flexoid/translators-map-go/internal/config"
	"github.com/flexoid/translators-map-go/internal/logging"
	"github.com/flexoid/translators-map-go/internal/services"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	logger := setupLogger()
	defer logger.Sync()

	kongCtx := kong.Parse(&config.CLI, kong.Configuration(kongyaml.Loader, "config/config.yml"))

	entClient, err := setupDatabase()
	if err != nil {
		logger.Fatalf("failed to setup database: %v", err)
	}

	switch kongCtx.Command() {
	case "server":
		startServer(entClient, logger, config.CLI.BindAddr)
	case "scraper":
		startScraper(entClient, logger)
	default:
		panic(kongCtx.Command())
	}
}

func startServer(entClient *ent.Client, logger *zap.SugaredLogger, bindAddr string) {
	logger.Info("Starting server")

	server := api.Server{EntDB: entClient, Logger: logger}

	err := server.Start(bindAddr)
	if err != nil {
		logger.Error(err)
	}
}

func startScraper(entClient *ent.Client, logger *zap.SugaredLogger) {
	for {
		logger.Info("Starting scraper")

		services.NewScraper(entClient, logger, config.CLI.MapsBackendAPIKey).Run()
		time.Sleep(24 * time.Hour)
	}
}

func setupDatabase() (*ent.Client, error) {
	db, err := sql.Open("postgres", config.CLI.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	drv := entsql.OpenDB("postgres", db)
	drvWithContext := dialect.DebugWithContext(drv, func(ctx context.Context, args ...interface{}) {
		logging.Ctx(ctx).Debug(args...)
	})

	client := ent.NewClient(ent.Driver(drvWithContext))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	return client, nil
}

func setupLogger() *zap.SugaredLogger {
	logCfg := zap.NewDevelopmentConfig()
	logCfg.Level.SetLevel(zap.DebugLevel)

	zapLogger := zap.Must(logCfg.Build())
	zap.ReplaceGlobals(zapLogger)

	return zapLogger.Sugar()
}
