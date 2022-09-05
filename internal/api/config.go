package api

import (
	"net/http"

	"github.com/flexoid/translators-map-go/internal/config"
)

type ConfigController struct {
	Server *Server
}

type Config struct {
	MapsJSAPIKey string `json:"maps_js_api_key"`
}

func (c *ConfigController) GetConfig(w http.ResponseWriter, r *http.Request) {
	config := Config{
		MapsJSAPIKey: config.CLI.MapsJSAPIKey,
	}

	err := encodeJSONResponse(config, 0, w, c.Server.Logger)
	if err != nil {
		c.Server.Logger.Infow("Failed to serialise activities", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
