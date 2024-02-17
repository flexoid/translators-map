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

type mockGeocoderInUK struct {
}

func (m *mockGeocoderInUK) Geocode(ctx context.Context, r *gmaps.GeocodingRequest) ([]gmaps.GeocodingResult, error) {
	return []gmaps.GeocodingResult{
		{
			AddressComponents: []gmaps.AddressComponent{
				{
					LongName:  "12 g5",
					ShortName: "12 g5",
					Types:     []string{"subpremise"},
				},
				{
					LongName:  "32",
					ShortName: "32",
					Types:     []string{"street_number"},
				},
				{
					LongName:  "Waddell Court",
					ShortName: "Waddell Ct",
					Types:     []string{"route"},
				},
				{
					LongName:  "Glasgow",
					ShortName: "Glasgow",
					Types:     []string{"postal_town"},
				},
				{
					LongName:  "Glasgow City",
					ShortName: "Glasgow City",
					Types:     []string{"administrative_area_level_2", "political"},
				},
				{
					LongName:  "Scotland",
					ShortName: "Scotland",
					Types:     []string{"administrative_area_level_1", "political"},
				},
				{
					LongName:  "United Kingdom",
					ShortName: "GB",
					Types:     []string{"country", "political"},
				},
				{
					LongName:  "G5 0PX",
					ShortName: "G5 0PX",
					Types:     []string{"postal_code"},
				},
			},
			FormattedAddress: "32 Waddell Ct, Glasgow G5 0PX, UK",
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

type mockGeocoderEmptyAddress struct {
}

func (m *mockGeocoderEmptyAddress) Geocode(ctx context.Context, r *gmaps.GeocodingRequest) ([]gmaps.GeocodingResult, error) {
	return []gmaps.GeocodingResult{
		{
			AddressComponents: []gmaps.AddressComponent{},
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

func TestGeocodingForAddressInPL(t *testing.T) {
	geocoder := NewGeocodingWithClient(&mockGeocoder{})
	result, err := geocoder.GeocodingForAddress(context.Background(), "Some address")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 52.2178902, result.Lat)
	assert.Equal(t, 20.9015052, result.Lng)
	assert.Equal(t, "Warszawa", result.City)
	assert.Equal(t, "Mazowieckie", result.AdministrativeArea)
	assert.Equal(t, "Poland", result.Country)
}

func TestGeocodingForAddressInUK(t *testing.T) {
	geocoder := NewGeocodingWithClient(&mockGeocoderInUK{})
	result, err := geocoder.GeocodingForAddress(context.Background(), "Some address")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Glasgow", result.City)
	assert.Equal(t, "Scotland", result.AdministrativeArea)
	assert.Equal(t, "United Kingdom", result.Country)
}

func TestGeocodingForAddressEmptyAddress(t *testing.T) {
	geocoder := NewGeocodingWithClient(&mockGeocoderEmptyAddress{})
	result, err := geocoder.GeocodingForAddress(context.Background(), "Some address")

	assert.NoError(t, err)
	assert.Equal(t, 52.2178902, result.Lat)
	assert.Equal(t, 20.9015052, result.Lng)
	assert.Equal(t, "", result.City)
	assert.Equal(t, "", result.AdministrativeArea)
	assert.Equal(t, "", result.Country)
}
