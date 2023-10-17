package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const metricNamespace = "translators_map"
const metricSubsystemScraper = "scraper"

type ScraperMetrics struct {
	TranslatorsScraped prometheus.Gauge
	TranslatorsCreated prometheus.Gauge
	TranslatorsUpdated prometheus.Gauge
	TranslatorsSkipped prometheus.Gauge
	LanguagesScraped   prometheus.Gauge
	CompletionTime     prometheus.Gauge
	SuccessTime        prometheus.Gauge
	Duration           prometheus.Gauge
}

func NewScraperMetrics() *ScraperMetrics {
	return &ScraperMetrics{
		TranslatorsScraped: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "translators_scraped",
			Help:      "Number of translators scraped during last run",
		}),
		TranslatorsCreated: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "translators_created",
			Help:      "Number of translators created during last run",
		}),
		TranslatorsUpdated: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "translators_updated",
			Help:      "Number of translators updated during last run",
		}),
		TranslatorsSkipped: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "translators_skipped",
			Help:      "Number of translators skipped as not changed during last run",
		}),
		LanguagesScraped: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "languages_scraped",
			Help:      "Number of languages scraped during last run",
		}),
		CompletionTime: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "completion_timestamp_seconds",
			Help:      "Time of last run of scraper",
		}),
		// Intentionally not registered with promauto, as we want to include it only if scraper succeeds.
		SuccessTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "success_timestamp_seconds",
			Help:      "Time of last successful run of scraper",
		}),
		Duration: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: metricNamespace,
			Subsystem: metricSubsystemScraper,
			Name:      "duration_seconds",
			Help:      "Duration of last run of scraper",
		}),
	}
}

func (m *ScraperMetrics) ResetBeforeRun() {
	m.TranslatorsScraped.Set(0)
	m.TranslatorsCreated.Set(0)
	m.TranslatorsUpdated.Set(0)
	m.TranslatorsSkipped.Set(0)
	m.LanguagesScraped.Set(0)
}
