package services

import (
	"context"
	"testing"

	"github.com/flexoid/translators-map-go/ent/enttest"
	"github.com/flexoid/translators-map-go/internal/metrics"
	"github.com/flexoid/translators-map-go/internal/scraper"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	gmaps "googlemaps.github.io/maps"
)

type mockGeocoder struct {
}

func (m *mockGeocoder) GetCoordinatesForAddress(ctx context.Context, address string) (*gmaps.GeocodingResult, error) {
	return &gmaps.GeocodingResult{
		Geometry: gmaps.AddressGeometry{
			Location: gmaps.LatLng{
				Lat: 123,
				Lng: 456,
			},
		},
	}, nil
}

func TestHandleTranslator(t *testing.T) {
	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer db.Close()

	logger := zap.NewNop().Sugar()
	metrics := metrics.NewScraperMetrics()

	s := NewScraper(db, logger, &mockGeocoder{}, metrics)

	trans := scraper.Translator{
		ID:       123,
		Name:     "John Doe",
		Address:  "123 Main St",
		Contacts: "555-555-5555",
		Language: scraper.Language{
			Name: "English",
			Code: 1,
		},
	}

	t.Run("new translator", func(t *testing.T) {
		model, err := s.handleTranslator(trans)
		assert.NoError(t, err)
		assert.NotNil(t, model)

		// Verify that the translator was saved to the database
		count, err := db.Translator.Query().Count(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 1, count)

		// Verify that the saved translator matches the original
		savedTrans, err := db.Translator.Query().First(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, trans.ID, savedTrans.ExternalID)
		assert.Equal(t, trans.Name, savedTrans.Name)
		assert.Equal(t, trans.Address, savedTrans.Address)
		assert.Equal(t, trans.Language.Name, savedTrans.Language)
	})

	t.Run("existing translator", func(t *testing.T) {
		_, err := s.handleTranslator(trans)
		assert.NoError(t, err)

		updatedTrans := scraper.Translator{
			ID:       123,
			Name:     "Jane Doe",
			Address:  "456 Main St",
			Contacts: "444-444-4444",
			Language: scraper.Language{
				Name: "English",
				Code: 1,
			},
		}

		model, err := s.handleTranslator(updatedTrans)
		assert.NoError(t, err)
		assert.NotNil(t, model)

		// Verify that no new translators were saved to the database
		count, err := db.Translator.Query().Count(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 1, count)

		// Verify that the saved translator matches the updated one
		savedTrans, err := db.Translator.Query().First(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, updatedTrans.ID, savedTrans.ExternalID)
		assert.Equal(t, updatedTrans.Name, savedTrans.Name)
		assert.Equal(t, updatedTrans.Address, savedTrans.Address)
		assert.Equal(t, updatedTrans.Language.Name, savedTrans.Language)
	})
}
