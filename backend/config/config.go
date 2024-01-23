package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIHost              string `envconfig:"API_HOST"`
	APIPort              string `envconfig:"API_PORT"               default:"8080"`
	CORSAllowOrigin      string `envconfig:"CORS_ALLOW_ORIGIN"      default:"*"`
	DbConnectionString   string `envconfig:"DB_CONNECTION_STRING"      default:"user=postgres password=shY74qI5zbbWhEMU dbname=postgres host=db.mlunrkdivnylgkxgudws.supabase.co"`
	CORSAllowCredentials string `envconfig:"CORS_ALLOW_CREDENTIALS" default:"true"`
	CORSAllowHeaders     string `envconfig:"CORS_ALLOW_HEADERS"     default:"*"`
	CORSAllowMethods     string `envconfig:"CORS_ALLOW_METHODS"     default:"GET, PUT, PATCH, POST, DELETE, OPTIONS"`
	VanguardUrl          string `envconfig:"VANGUARD_URL"     default:"https://www.nl.vanguard/gpx/graphql"`
	ISharesUrl           string `envconfig:"ISHARES_URL"     default:"https://www.ishares.com"`
	// NL
	ISharesLocalePath            string `envconfig:"ISHARES_LOCALE_PATH" default:"/nl/professionele-belegger/nl"`
	ISharesSearchPage            string `envconfig:"ISHARES_SEARCH_PAGE" default:"/producten/etf-investments?switchLocale=y&siteEntryPassthrough=true#/?productView=etf&"`
	ISharesProductRegex          string `envconfig:"ISHARES_PRODUCT_REGEX" default:"/producten/([0-9]*)"`
	ISharesHoldingCallIdentifier string `envconfig:"ISHARES_HOLDING_CALL_IDENTIFIER" default:"1497735778849"`
	// UK
	//ISharesLocalePath            string `envconfig:"ISHARES_LOCALE_PATH" default:"/uk/professional/en"`
	//ISharesSearchPage            string `envconfig:"ISHARES_SEARCH_PAGE" default:"/products/etf-investments?switchLocale=y&siteEntryPassthrough=true#/?productView=etf&"`
	//ISharesProductRegex          string `envconfig:"ISHARES_PRODUCT_REGEX" default:"/products/([0-9]*)"`
	//ISharesHoldingCallIdentifier string `envconfig:"ISHARES_HOLDING_CALL_IDENTIFIER" default:"1506575576011"`
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
