package hackerone

import (
	"net/http"
	"os"
	"testing"

	"github.com/ayuxsec-org/scopex/pkg/config"
	"github.com/joho/godotenv"
)

var scraper *Scraper

func TestMain(m *testing.M) {
	_ = godotenv.Load("../../../.env")

	scraper = NewScraper(config.HackerOneCfg{
		Creds: config.HackerOneApiCreds{
			UserName: os.Getenv("HACKERONE_USER_NAME"),
			Password: os.Getenv("HACKERONE_API_KEY"),
		},
		RateLimitPerMin: 600,
	}, http.DefaultClient)

	code := m.Run()
	os.Exit(code)
}
