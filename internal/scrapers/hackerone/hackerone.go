package hackerone

import (
	"net/http"

	"github.com/ayuxsec-org/scopex/internal/request"
	"github.com/ayuxsec-org/scopex/pkg/config"
)

// Scraper handles HackerOne API interactions.
//
// Scraper must be created using NewScraper.
type Scraper struct {
	r *request.Request
}

// NewScraper constructs and initializes a HackerOne scraper.
func NewScraper(cfg config.HackerOneCfg, client *http.Client) *Scraper {
	if client == nil {
		client = http.DefaultClient
	}
	return &Scraper{
		r: &request.Request{
			UserName: cfg.Creds.UserName,
			Password: cfg.Creds.Password,
			Client:   client,
		},
	}
}
