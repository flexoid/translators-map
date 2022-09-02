package maps

import (
	"context"
	"fmt"

	"googlemaps.github.io/maps"
)

type Geocoding struct {
	client *maps.Client
}

func NewGeocoding(apiKey string) (*Geocoding, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &Geocoding{
		client: c,
	}, nil
}

func (g *Geocoding) GetCoordinatesForAddress(ctx context.Context, address string) (*maps.GeocodingResult, error) {
	r := &maps.GeocodingRequest{
		Address: address,
	}

	result, err := g.client.Geocode(ctx, r)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("cannot find coordinates for address: %s", address)
	}

	return &result[0], nil
}
