package api

import (
	"net/http"

	"github.com/flexoid/translators-map-go/ent/translator"
)

type TranslatorController struct {
	Server *Server
}

type Translator struct {
	DetailsURL string             `json:"details_url"`
	Location   TranslatorLocation `json:"location"`
}

type TranslatorLocation struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

func (c *TranslatorController) GetTranslators(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		encodeJSONResponse(struct {
			Error string `json:"error"`
		}{Error: "lang parameter is required"}, http.StatusBadRequest, w, c.Server.Logger)

		return
	}

	translators := []Translator{}

	dbTranslators, err := c.Server.EntDB.Translator.Query().
		Where(translator.Language(lang)).All(r.Context())
	if err != nil {
		c.Server.Logger.Errorw("Failed to get translators", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, dbTranslator := range dbTranslators {
		translators = append(translators, Translator{
			Location: TranslatorLocation{
				Latitude:  dbTranslator.Latitude,
				Longitude: dbTranslator.Longitude,
			},
			DetailsURL: dbTranslator.DetailsURL,
		})
	}

	err = encodeJSONResponse(translators, 0, w, c.Server.Logger)
	if err != nil {
		c.Server.Logger.Infow("Failed to serialise activities", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
