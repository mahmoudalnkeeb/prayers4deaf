package prayer

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/mahmoudalnkeeb/prayers4deaf/geolocation"
	"github.com/mahmoudalnkeeb/prayers4deaf/utils"
)

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		Timings struct {
			Fajr    string `json:"Fajr"`
			Dhuhr   string `json:"Dhuhr"`
			Asr     string `json:"Asr"`
			Maghrib string `json:"Maghrib"`
			Isha    string `json:"Isha"`
		} `json:"timings"`
	} `json:"data"`
}

// Embadded integration:
// could be updated by assigning pin for leds here
// example machine.Pin(0)
type Prayer struct {
	Time time.Time
	Name string
	LED  string // type => machine.Pin
}

func (p *Prayer) AssignLed() {
	p.LED = "led"
}

// GetPrayers fetches prayer times from the external API and returns them as a slice of Prayers
func GetPrayers(geo *geolocation.GeoLocation, logger *slog.Logger) ([]Prayer, error) {
	date := time.Now().Format(time.DateOnly)
	baseUrl := fmt.Sprintf("http://api.aladhan.com/v1/timingsByCity/%s", date)

	logger.Info("Fetching prayer times from http://api.aladhan.com/v1/timingsByCity for", "date", date)

	params := []utils.Param{
		{Name: "latitude", Value: geo.Latitude},
		{Name: "longitude", Value: geo.Longitude},
		{Name: "x7xapikey", Value: os.Getenv("X7X_API_KEY")},
		{Name: "city", Value: geo.City},
		{Name: "country", Value: geo.CountryCode},
		{Name: "method", Value: "5"},
	}

	resp, err := utils.CreateGetRequest(baseUrl, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get prayer times: %w", err)
	}
	defer resp.Body.Close()

	var response Response
	if err := utils.DecodeJsonResponse(resp.Body, &response); err != nil {
		return nil, fmt.Errorf("failed to decode prayer times response: %w", err)
	}

	if response.Code != 200 {
		return nil, fmt.Errorf("unexpected API response status: %s", response.Status)
	}

	// Get prayer times from response then parse it to Time and apply DST
	prayers := []Prayer{
		{Name: "Fajr", Time: parseTime(response.Data.Timings.Fajr, logger, true)},
		{Name: "Dhuhr", Time: parseTime(response.Data.Timings.Dhuhr, logger, true)},
		{Name: "Asr", Time: parseTime(response.Data.Timings.Asr, logger, true)},
		{Name: "Maghrib", Time: parseTime(response.Data.Timings.Maghrib, logger, true)},
		{Name: "Isha", Time: parseTime(response.Data.Timings.Isha, logger, true)},
	}

	return prayers, nil
}

// parseTime convert time string to Time type and apply DST if needed
func parseTime(prayerTime string, logger *slog.Logger, dst bool) time.Time {
	parsedTime, err := time.Parse("15:04", prayerTime)
	if err != nil {
		logger.Error("Error while parsing time", "err", err)
		return time.Time{} // return zero time if parsing fails
	}
	if dst {
		return parsedTime.Add(60 * time.Minute)
	}
	return parsedTime
}

func GetNextPrayer(prayers []Prayer, currentTime string, logger *slog.Logger) (string, string, error) {
	const timeFormat = "15:04" // HH:MM format
	current := parseTime(currentTime, logger, false)

	for _, prayer := range prayers {
		if prayer.Time.After(current) {
			return prayer.Name, prayer.Time.Format(timeFormat), nil
		}
	}

	return "Fajr", prayers[0].Time.Format(timeFormat), nil
}

// GetCurrentPrayer returns the ongoing prayer if the current time is within
// the prayer time and the prayer duration.
//
// prayerDuration represents the duration in minutes for which the current
// prayer is considered ongoing. Consider renaming it to something more
// descriptive, such as `ongoingPrayerWindow` or `prayerTimeFrame`.
func GetCurrentPrayer(prayers []Prayer, currentTime string, prayerDuration time.Duration, logger *slog.Logger) (string, string, error) {
	const timeFormat = "15:04" // HH:MM format
	current := parseTime(currentTime, logger, false)

	for _, prayer := range prayers {
		timeWindow := prayer.Time.Add(prayerDuration * time.Minute)

		if current.After(prayer.Time) && current.Before(timeWindow) {
			return prayer.Name, prayer.Time.Format(timeFormat), nil
		}
	}

	return "", "", nil
}
