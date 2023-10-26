package configuration

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jftuga/geodist"
)

// Config is the type used for all user configuration.
// All parameters can be set using ENV variables.
// The comments below are structured as following:
// ENV_VARIABLE_NAME DEFAULT_VALUE
type Config struct {
	// Latitude and Longitude coordinates of the location you want to use.
	// LOCATION_LATITUDE 51.17348
	// LOCATION_LONGITUDE 5.45921
	Location geodist.Coord

	// Maximum range in kilometers from the location that you want aircraft to be spotted.
	// Note that this is an approximation due to roundings.
	// MAX_RANGE_KILOMETERS 30
	MaxRangeKilometers int

	// A comma seperated list of types that you want to spot
	// If not set, 'ALL' will be used, which will disable the filter and show all aircraft within range.
	// Full list can be found at https://www.icao.int/publications/doc8643/pages/search.aspx in 'Type Designator' column.
	// AIRCRAFT_TYPES ALL
	// EXAMPLES
	// AIRCRAFT_TYPES F16,F35
	AircraftTypes []string

	// Maximum amount of aircraft to show in a single slack message.
	// Note that a single slack message only supports up to 50 'blocks' and each aircraft that we display has multiple blocks.
	// MAX_AIRCRAFT_SLACK_MESSAGE 8
	MaxAircraftSlackMessage int

	// Webhook used to send notifications to Slack. If not set, no messages will be sent to Slack.
	// SLACK_WEBHOOK_URL ""
	SlackWebHookURL string

	// Webhook used to send notifications to Discord. If not set, no messages will be sent to Discord.
	// DISCORD_WEBHOOK_URL ""
	DiscordWebHookURL string

	// Discord notifications use an embed color based on the alitute of the aircraft.
	// DISCORD_COLOR_ALTITUDE "true"
	DiscordColorAltitude string

	// Interval in seconds between fetching aircraft, minimum is 60 due to API rate limiting.
	// FETCH_INTERVAL 60
	FetchInterval int

	// Token to authenticate with the gotify server.
	// GOTIFY_TOKEN ""
	GotifyToken string

	// URL of the gotify server.
	// GOTIFY_URL ""
	GotifyURL string

	// Port where metrics will be exposed on
	// METRICS_PORT "7070"
	MetricsPort string
}

// Environment variable names
const (
	SlackWebhookURL          = "SLACK_WEBHOOK_URL"
	DiscordWebhookURL        = "DISCORD_WEBHOOK_URL"
	DiscordColorAltitude     = "DISCORD_COLOR_ALTITUDE"
	LocationLatitude         = "LOCATION_LATITUDE"
	LocationLongitude        = "LOCATION_LONGITUDE"
	MaxRangeKilometers       = "MAX_RANGE_KILOMETERS"
	MaxAircrfaftSlackMessage = "MAX_AIRCRAFT_SLACK_MESSAGE"
	AircraftTypes            = "AIRCRAFT_TYPES"
	FetchInterval            = "FETCH_INTERVAL"
	GotifyURL                = "GOTIFY_URL"
	GotifyToken              = "GOTIFY_TOKEN"
	MetricsPort              = "METRICS_PORT"
)

// getEnvVariable looks up a specified environment variable, if not set the specified default is used
func getEnvVariable(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

// GetConfig attempts to read the configuration via environment variables and uses a default if the environment variable is not set
func GetConfig() (config Config, err error) {
	defaultFetchInterval := 60

	config.GotifyToken = getEnvVariable(GotifyToken, "")
	config.GotifyURL = getEnvVariable(GotifyURL, "")
	config.SlackWebHookURL = getEnvVariable(SlackWebhookURL, "")
	config.DiscordWebHookURL = getEnvVariable(DiscordWebhookURL, "")
	config.DiscordColorAltitude = getEnvVariable(DiscordColorAltitude, "true")
	config.MetricsPort = getEnvVariable(MetricsPort, "7070")
	config.FetchInterval, err = strconv.Atoi(getEnvVariable(FetchInterval, strconv.Itoa(defaultFetchInterval)))
	if err != nil || config.FetchInterval < 60 {
		log.Printf("Fetch interval of %ds detected. You might hit rate limits, consider using the default of %ds instead.", config.FetchInterval, defaultFetchInterval)
	}

	config.Location.Lat, err = strconv.ParseFloat(getEnvVariable(LocationLatitude, "51.17348"), 64)
	if err != nil {
		return Config{}, err
	}

	config.Location.Lon, err = strconv.ParseFloat(getEnvVariable(LocationLongitude, "5.45921"), 64)
	if err != nil {
		return Config{}, err
	}

	config.MaxRangeKilometers, err = strconv.Atoi(getEnvVariable(MaxRangeKilometers, "30"))
	if err != nil {
		return Config{}, err
	}

	config.MaxAircraftSlackMessage, err = strconv.Atoi(getEnvVariable(MaxAircrfaftSlackMessage, "8"))
	if err != nil {
		return Config{}, err
	}

	config.AircraftTypes = strings.Split(strings.ToUpper(strings.ReplaceAll(getEnvVariable(AircraftTypes, "ALL"), " ", "")), ",")
	return config, nil
}
