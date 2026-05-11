package main

import "github.com/ayuxsec-org/scopex/pkg/config"

func NewCmdi() *Cmdi {
	return &Cmdi{}
}

type Cmdi struct {
	Config    config.Config
	Hackerone HackerOneCmdi
}

type HackerOneCmdi struct {
	ScrapeWildCards  bool
	ScrapeDomains    bool
	ScrapeSourceCode bool
	VDPOnly          bool
}
