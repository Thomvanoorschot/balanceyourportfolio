package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiHost              string `envconfig:"API_HOST"`
	ApiPort              string `envconfig:"API_PORT"               default:"8080"`
	CorsAllowOrigin      string `envconfig:"CORS_ALLOW_ORIGIN"      default:"*"`
	DbConnectionString   string `envconfig:"DB_CONNECTION_STRING"      default:"user=postgres password=shY74qI5zbbWhEMU dbname=postgres host=db.mlunrkdivnylgkxgudws.supabase.co"`
	CorsAllowCredentials string `envconfig:"CORS_ALLOW_CREDENTIALS" default:"true"`
	CorsAllowHeaders     string `envconfig:"CORS_ALLOW_HEADERS"     default:"*"`
	CorsAllowMethods     string `envconfig:"CORS_ALLOW_METHODS"     default:"GET, PUT, PATCH, POST, DELETE, OPTIONS"`
	VanguardUrl          string `envconfig:"VANGUARD_URL"     default:"https://www.nl.vanguard/gpx/graphql"`
	ISharesUrl           string `envconfig:"ISHARES_URL"     default:"https://www.ishares.com"`
}

var config *Config
var once sync.Once

// Load reads config file and ENV variables if set.
func Load() *Config {
	once.Do(func() {
		load()

	})

	return config
}

func load() {
	config = new(Config)
	if err := envconfig.Process("", config); err != nil {
		log.Fatal(err)
	}
}
