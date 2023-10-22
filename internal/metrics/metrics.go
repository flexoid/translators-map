package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const metricNamespace = "translators_map"
const metricSubsystemScraper = "scraper"

type ScraperMetrics struct {
	TranslatorsScraped prometheus.Gauge
	TranslatorsCreated prometheus.Gauge
	TranslatorsUpdated prometheus.Gauge
	TranslatorsDeleted prometheus.Gauge
	TranslatorsFailed  prometheus.Gauge
	LanguagesScraped   prometheus.Gauge
	CompletionTime     prometheus.Gauge
	SuccessTime        prometheus.Gauge
	Duration           prometheus.Gauge
}

func NewScraperMetrics(registry prometheus.Registerer) *ScraperMetrics {
	m := ScraperMetrics{
		TranslatorsScraped: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "translators_scraped",
			Help:      "Number of translators scraped during last run",
		}),
		TranslatorsCreated: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "translators_created",
			Help:      "Number of translators created during last run",
		}),
		TranslatorsUpdated: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "translators_updated",
			Help:      "Number of translators updated during last run",
		}),
		TranslatorsDeleted: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "translators_deleted",
			Help:      "Number of translators deleted during last run",
		}),
		TranslatorsFailed: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "translators_failed",
			Help:      "Number of translators failed to process during last run",
		}),
		LanguagesScraped: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "languages_scraped",
			Help:      "Number of languages scraped during last run",
		}),
		CompletionTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "completion_timestamp_seconds",
			Help:      "Time of last run of scraper",
		}),
		SuccessTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "success_timestamp_seconds",
			Help:      "Time of last successful run of scraper",
		}),
		Duration: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "duration_seconds",
			Help:      "Duration of last run of scraper",
		}),
	}

	// Register metrics with prometheus.
	// Intentionally not registered `SuccessTime`, as we want to include it only if scraper succeeds.
	registry.MustRegister(
		m.TranslatorsScraped,
		m.TranslatorsCreated,
		m.TranslatorsUpdated,
		m.TranslatorsDeleted,
		m.TranslatorsFailed,
		m.LanguagesScraped,
		m.CompletionTime,
		m.Duration,
	)

	return &m
}

func (m *ScraperMetrics) ResetBeforeRun() {
	m.TranslatorsScraped.Set(0)
	m.TranslatorsCreated.Set(0)
	m.TranslatorsUpdated.Set(0)
	m.TranslatorsDeleted.Set(0)
	m.TranslatorsFailed.Set(0)
	m.LanguagesScraped.Set(0)
}
