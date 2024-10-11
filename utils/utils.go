package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Param struct holds query parameters
type Param struct {
	Name  string
	Value string
}

// HTTP client with timeout
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// CreateGetRequest creates an HTTP GET request with query parameters
func CreateGetRequest(url string, params []Param) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	setParams(req, params)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp, nil
}

// setParams adds query parameters to the HTTP request
func setParams(req *http.Request, params []Param) {
	q := req.URL.Query()
	for _, param := range params {
		q.Set(param.Name, param.Value)
	}
	req.URL.RawQuery = q.Encode()
}

// DecodeJsonResponse decodes JSON response to the target structure
func DecodeJsonResponse[T any](body io.Reader, target *T) error {
	if err := json.NewDecoder(body).Decode(target); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}
	return nil
}

// GetCurrentTime returns the current time in HH:MM format
func GetCurrentTime() string {
	h, m, _ := time.Now().Clock()
	return fmt.Sprintf("%02d:%02d", h, m)
}

// FilterPrayerTimes filters the response to return only the five main prayers
func FilterPrayerTimes(timings map[string]string) map[string]string {
	mainPrayers := []string{"Fajr", "Dhuhr", "Asr", "Maghrib", "Isha"}
	filteredPrayers := make(map[string]string)

	for _, prayer := range mainPrayers {
		if time, exists := timings[prayer]; exists {
			filteredPrayers[prayer] = time
		}
	}

	return filteredPrayers
}
