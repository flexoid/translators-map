package services

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/flexoid/translators-map-go/ent"
	"github.com/flexoid/translators-map-go/ent/enttest"
	"github.com/flexoid/translators-map-go/ent/translator"
	"github.com/flexoid/translators-map-go/internal/maps"
	"github.com/flexoid/translators-map-go/internal/metrics"
	"github.com/flexoid/translators-map-go/internal/scraper"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type mockGeocoder struct {
}

func (m *mockGeocoder) GeocodingForAddress(ctx context.Context, address string) (*maps.Result, error) {
	switch address {
	case "123 Main St":
		return &maps.Result{
			Lat:                123,
			Lng:                456,
			City:               "Warszawa",
			AdministrativeArea: "Mazowieckie",
			Country:            "Poland",
		}, nil
	case "456 Main St":
		return &maps.Result{
			Lat:                456,
			Lng:                789,
			City:               "Kraków",
			AdministrativeArea: "Małopolskie",
			Country:            "Poland",
		}, nil
	default:
		return nil, fmt.Errorf("unknown address: %s", address)
	}
}

func TestHandleTranslator(t *testing.T) {
	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer db.Close()

	logger := zap.NewNop().Sugar()
	metrics := metrics.NewScraperMetrics(prometheus.NewRegistry())

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
		assert.Equal(t, "Warszawa", savedTrans.City)
		assert.Equal(t, "Mazowieckie", savedTrans.AdministrativeArea)
		assert.Equal(t, "Poland", savedTrans.Country)
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
		assert.Equal(t, "Kraków", savedTrans.City)
		assert.Equal(t, "Małopolskie", savedTrans.AdministrativeArea)
		assert.Equal(t, "Poland", savedTrans.Country)
	})
}

func TestDeleteOldTranslators(t *testing.T) {
	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer db.Close()

	s := NewScraper(db, zap.NewNop().Sugar(), &mockGeocoder{},
		metrics.NewScraperMetrics(prometheus.NewRegistry()))

	externalIDs := []int{10, 11, 12, 13, 14}
	_, err := db.Translator.MapCreateBulk(externalIDs, func(tc *ent.TranslatorCreate, i int) {
		tc.SetExternalID(externalIDs[i]).
			SetLanguage("English").
			SetName("John Doe").
			SetAddress("123 Main St").
			SetAddressSha([]byte("")).
			SetDetailsURL("https://example.com").
			SetCity("Warszawa").
			SetAdministrativeArea("Mazowieckie").
			SetCountry("Poland")
	}).Save(context.Background())
	require.NoError(t, err)

	translators, _ := db.Translator.Query().Order(ent.Asc(translator.FieldExternalID)).
		All(context.Background())

	translators[0].Update().SetUpdatedAt(time.Now().Add(-DeleteOldTranslatorsInterval/3 - 5*time.Minute)).
		Save(context.Background())
	translators[1].Update().SetUpdatedAt(time.Now().Add(-DeleteOldTranslatorsInterval*2 - 5*time.Minute)).
		Save(context.Background())
	translators[2].Update().SetUpdatedAt(time.Now().Add(-DeleteOldTranslatorsInterval - 5*time.Minute)).
		Save(context.Background())
	translators[4].Update().SetUpdatedAt(time.Now().Add(-DeleteOldTranslatorsInterval/2 - 5*time.Minute)).
		Save(context.Background())

	s.deleteOldTranslators()

	translatorsLeft, err := db.Translator.Query().Order(ent.Asc(translator.FieldExternalID)).
		All(context.Background())

	require.NoError(t, err)
	require.Len(t, translatorsLeft, 3)

	assert.Equal(t, 10, translatorsLeft[0].ExternalID)
	assert.Equal(t, 13, translatorsLeft[1].ExternalID)
	assert.Equal(t, 14, translatorsLeft[2].ExternalID)
}
