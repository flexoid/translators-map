package maps

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	gmaps "googlemaps.github.io/maps"
)

type mockGeocoder struct {
}

func (m *mockGeocoder) Geocode(ctx context.Context, r *gmaps.GeocodingRequest) ([]gmaps.GeocodingResult, error) {
	return []gmaps.GeocodingResult{
		{
			AddressComponents: []gmaps.AddressComponent{
				{
					LongName:  "1",
					ShortName: "1",
					Types:     []string{"subpremise"},
				},
				{
					LongName:  "2A",
					ShortName: "2A",
					Types:     []string{"street_number"},
				},
				{
					LongName:  "Szeligowska",
					ShortName: "Szeligowska",
					Types:     []string{"route"},
				},
				{
					LongName:  "Bemowo",
					ShortName: "Bemowo",
					Types:     []string{"political", "sublocality", "sublocality_level_1"},
				},
				{
					LongName:  "Warszawa",
					ShortName: "Warszawa",
					Types:     []string{"locality", "political"},
				},
				{
					LongName:  "Warszawa",
					ShortName: "Warszawa",
					Types:     []string{"administrative_area_level_2", "political"},
				},
				{
					LongName:  "Mazowieckie",
					ShortName: "Mazowieckie",
					Types:     []string{"administrative_area_level_1", "political"},
				},
				{
					LongName:  "Poland",
					ShortName: "PL",
					Types:     []string{"country", "political"},
				},
				{
					LongName:  "01-319",
					ShortName: "01-319",
					Types:     []string{"postal_code"},
				},
			},
			FormattedAddress: "Szeligowska 123, 01-319 Warszawa, Poland",
			Geometry: gmaps.AddressGeometry{
				Location: gmaps.LatLng{
					Lat: 52.2178902,
					Lng: 20.9015052,
				},
				Types: []string{},
			},
		},
	}, nil
}

func TestGeocodingForAddress(t *testing.T) {
	geocoder := NewGeocodingWithClient(&mockGeocoder{})
	result, err := geocoder.GeocodingForAddress(context.Background(), "Some address")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Warszawa", result.City)
	assert.Equal(t, "Mazowieckie", result.AdministrativeArea)
	assert.Equal(t, "Poland", result.Country)
	assert.Equal(t, 52.2178902, result.Lat)
	assert.Equal(t, 20.9015052, result.Lng)
}
