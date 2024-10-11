package main

import (
	"log/slog"
	"time"

	"github.com/mahmoudalnkeeb/prayers4deaf/geolocation"
	"github.com/mahmoudalnkeeb/prayers4deaf/prayer"
	"github.com/mahmoudalnkeeb/prayers4deaf/utils"

	"github.com/joho/godotenv"
)

var logger = slog.Default()

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	// Fetch geo location
	geo, err := geolocation.GetGeoLocation(logger)
	if err != nil {
		logger.Error("Error getting geo location", "error", err)
		return
	}

	// Fetch prayer times
	prayers, err := prayer.GetPrayers(geo, logger)
	if err != nil {
		logger.Error("Error fetching prayers", "error", err)
		return
	}
	currentTime := utils.GetCurrentTime()
	playerDuration := time.Duration(60)
	currentPrayerName, currentPrayerTime, err := prayer.GetCurrentPrayer(prayers, currentTime, playerDuration, logger)
	if err != nil {
		logger.Error("Error getting current prayer", "error", err)
		return
	}

	nextPrayerName, nextPrayerTime, err := prayer.GetNextPrayer(prayers, currentTime, logger)
	if err != nil {
		logger.Error("Error getting next prayer", "error", err)
		return
	}

	// Log current and next prayers
	logger.Info("Current ongoing prayer", "name", currentPrayerName, "prayer_time", currentPrayerTime, "current_time", currentTime)
	logger.Info("Next upcoming prayer", "name", nextPrayerName, "time", nextPrayerTime, "current_time", currentTime)
}
