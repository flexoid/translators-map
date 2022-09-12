package config

var CLI struct {
	BindAddr          string `required:"" env:"BIND_ADDR" help:"Bind address with port" default:"localhost:4000"`
	DatabaseURL       string `required:"" env:"DATABASE_URL" help:"Postgres database URL."`
	MapsBackendAPIKey string `required:"" env:"MAPS_BACKEND_API_KEY" help:"Key for Google Maps backend API."`
	MapsJSAPIKey      string `required:"" env:"MAPS_JS_API_KEY" help:"Key for Google Maps JavaScript API."`
	GoogleAnalyticsID string `env:"GOOGLE_ANALYTICS_ID" help:"Google Analytics ID."`

	Scraper struct {
	} `cmd:"" help:"Run scraper."`

	Server struct {
	} `cmd:"" help:"Run server."`
}
