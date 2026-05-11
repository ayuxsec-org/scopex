package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ayuxsec-org/log"
	"github.com/ayuxsec-org/scopex/internal/scrapers/hackerone"
	"github.com/spf13/cobra"
)

func (cmdi *Cmdi) NewHackerOneCmd() *cobra.Command {
	h1Cmd := &cobra.Command{
		Use:   "h1",
		Short: "Scrape scopes from hackerone",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return cmdi.validateH1()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdi.RunH1()
		},
	}

	h1Cmd.Flags().StringVarP(&cmdi.Config.Hackerone.Creds.UserName, "username", "u", "", "Username of your hackerone account")
	h1Cmd.Flags().StringVarP(&cmdi.Config.Hackerone.Creds.Password, "token", "t", "", "API token of your hackerone account")
	h1Cmd.Flags().BoolVarP(&cmdi.Hackerone.ScrapeDomains, "domains", "d", false, "Scrape all domains")
	h1Cmd.Flags().BoolVarP(&cmdi.Hackerone.ScrapeWildCards, "wc", "w", false, "Scrape all wildcards")
	h1Cmd.Flags().BoolVarP(&cmdi.Hackerone.ScrapeSourceCode, "source-code", "s", false, "find source codes")
	h1Cmd.Flags().BoolVar(&cmdi.Hackerone.VDPOnly, "vdp", false, "List VDP programs only")

	return h1Cmd
}

func (cmdi *Cmdi) validateH1() error {
	if cmdi.Config.Hackerone.Creds.UserName == "" || cmdi.Config.Hackerone.Creds.Password == "" {
		return errors.New("either hackerone username or api token is missing")
	}
	return nil
}

func (cmdi *Cmdi) RunH1() error {
	if !cmdi.Hackerone.ScrapeDomains || !cmdi.Hackerone.ScrapeWildCards || !cmdi.Hackerone.ScrapeSourceCode {
		log.Error("no argument provided for hackerone scraper. Run -h or --help to get more info")
		return nil
	}
	scraper := hackerone.NewScraper(cmdi.Config.Hackerone, http.DefaultClient)
	handles, err := scraper.GetAllHandles()
	if err != nil {
		return fmt.Errorf("'GetAllHandles': %v", err)
	}
	for _, h := range handles {
		handleScopes, err := scraper.GetHandleScopes(h.Handle)
		if err != nil {
			return fmt.Errorf("'GetHandleScopes': %v", err)
		}
		for _, scope := range handleScopes {
			if cmdi.Hackerone.ScrapeWildCards {
				if scope.EligibleForBounty && !cmdi.Hackerone.VDPOnly {
					// todo: add hardcoded values to a constant
					if scope.Type == "WILDCARD" {
						fmt.Println(scope.ID)
					}
				}

			}
			if cmdi.Hackerone.ScrapeDomains {
				if scope.EligibleForBounty && !cmdi.Hackerone.VDPOnly {
					if scope.Type == "URL" {
						fmt.Println(scope.ID)
					}
				}
			}
			if cmdi.Hackerone.ScrapeSourceCode {
				if scope.EligibleForBounty && !cmdi.Hackerone.VDPOnly {
					if scope.Type == "SOURCE_CODE" {
						fmt.Println(scope.ID)
					}
				}
			}
		}
	}
	return nil
}
