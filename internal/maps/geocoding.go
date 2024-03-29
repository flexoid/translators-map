package maps

import (
	"context"
	"fmt"

	"github.com/flexoid/translators-map-go/internal/logging"
	gmaps "googlemaps.github.io/maps"
)

type Geocoding struct {
	client MapsGeocoder
}

type MapsGeocoder interface {
	Geocode(ctx context.Context, r *gmaps.GeocodingRequest) ([]gmaps.GeocodingResult, error)
}

type Result struct {
	Lat                float64
	Lng                float64
	Country            string
	AdministrativeArea string
	City               string
}

const localityType = "locality"
const postalTownType = "postal_town"
const administrativeAreaType = "administrative_area_level_1"
const countryType = "country"

func NewGeocoding(apiKey string) (*Geocoding, error) {
	client, err := gmaps.NewClient(gmaps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return NewGeocodingWithClient(client), nil
}

func NewGeocodingWithClient(client MapsGeocoder) *Geocoding {
	return &Geocoding{
		client: client,
	}
}

func (g *Geocoding) GeocodingForAddress(ctx context.Context, address string) (*Result, error) {
	r := &gmaps.GeocodingRequest{
		Address: address,
	}

	result, err := g.client.Geocode(ctx, r)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("cannot find coordinates for address: %s", address)
	}

	logging.Ctx(ctx).Debugf("Geocoding result: %+v", result[0])

	city := g.extractComponent(result[0], localityType)
	if city == "" {
		city = g.extractComponent(result[0], postalTownType)
	}

	administrativeArea := g.extractComponent(result[0], administrativeAreaType)
	country := g.extractComponent(result[0], countryType)

	return &Result{
		Lat:                result[0].Geometry.Location.Lat,
		Lng:                result[0].Geometry.Location.Lng,
		Country:            country,
		AdministrativeArea: administrativeArea,
		City:               city,
	}, nil
}

func (g *Geocoding) extractComponent(result gmaps.GeocodingResult, componentType string) string {
	for _, component := range result.AddressComponents {
		for _, t := range component.Types {
			if t == componentType {
				return component.LongName
			}
		}
	}

	return ""
}
