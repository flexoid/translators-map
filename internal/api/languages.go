package api

import (
	"net/http"

	"github.com/flexoid/translators-map-go/ent"
	"github.com/flexoid/translators-map-go/ent/translator"
)

type LanguageController struct {
	Server *Server
}

type Language struct {
	Language string `json:"language"`
}

func (c *LanguageController) GetLanguages(w http.ResponseWriter, r *http.Request) {
	languages := []Language{}

	dbLangs, err := c.Server.EntDB.Translator.Query().Order(ent.Asc(translator.FieldLanguage)).
		GroupBy(translator.FieldLanguage).Strings(r.Context())
	if err != nil {
		c.Server.Logger.Errorw("Failed to get languages from database", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, dbLang := range dbLangs {
		languages = append(languages, Language{
			Language: dbLang,
		})
	}

	err = encodeJSONResponse(languages, 0, w, c.Server.Logger)
	if err != nil {
		c.Server.Logger.Errorw("Failed to serialise languages", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
