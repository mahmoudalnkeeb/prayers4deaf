package geolocation

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/mahmoudalnkeeb/prayers4deaf/utils"
)

const apiUrl = "https://api.ipgeolocation.io/ipgeo"

type GeoLocation struct {
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	CountryCode string `json:"country_code2"`
	City        string `json:"city"`
}

// GetGeoLocation fetches geographical location from an external API
func GetGeoLocation(logger *slog.Logger) (*GeoLocation, error) {
	apiKey := os.Getenv("IPGEO_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("IPGEO_API_KEY not set")
	}

	logger.Info("Fetching GeoLocation info from ipgeolocation.io/ipgeo...")

	params := []utils.Param{
		{Name: "apiKey", Value: apiKey},
	}

	resp, err := utils.CreateGetRequest(apiUrl, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get geo location: %w", err)
	}
	defer resp.Body.Close()

	var location GeoLocation
	if err := utils.DecodeJsonResponse(resp.Body, &location); err != nil {
		return nil, fmt.Errorf("failed to decode geo location response: %w", err)
	}

	return &location, nil
}
