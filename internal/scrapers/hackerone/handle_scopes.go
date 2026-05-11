package hackerone

import (
	"fmt"
	"io"
	"strings"

	"github.com/ayuxsec-org/scopex/internal/utils"
	"github.com/tidwall/gjson"
)

// Scope represents a scope of a program handle
type Scope struct {
	// ID is the identifier of the scope
	ID string `json:"asset_identifier"`
	// EligibleForBounty is whether the scope is eligible for bounty
	EligibleForBounty bool `json:"eligible_for_bounty"`
	// Type is the type of the scope
	Type string `json:"asset_type"`
}

// GetAllHandlesAndTheirScopes fetches all handles and their scopes
func (s *Scraper) GetAllHandlesAndTheirScopes() ([]Scope, error) {
	var scopes []Scope
	handles, err := s.GetAllHandles()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrGettingProgramHandles, err)
	}
	for _, h := range handles {
		handleScope, err := s.GetHandleScopes(h.Handle)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrGettingProgramScopes, err)
		}
		scopes = append(scopes, handleScope...)
	}
	return scopes, nil
}

// GetHandleScopes fetches the scope for a given handle
//
// handle is the program's handle name
func (s *Scraper) GetHandleScopes(handle string) ([]Scope, error) {
	resp, err := s.r.Get(strings.ReplaceAll(ProgramHandleScopesAPIUrl, "{handle}", handle))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	respBody := string(utils.Must(io.ReadAll(resp.Body)))
	return extractStructuredScopes(respBody), nil
}

// extractStructuredScopes extracts structured scopes from the JSON string
func extractStructuredScopes(jsonStr string) []Scope {
	results := gjson.Get(jsonStr, "data.#.attributes")
	var scopes []Scope
	for _, item := range results.Array() {
		s := Scope{
			ID:                item.Get("asset_identifier").String(),
			EligibleForBounty: item.Get("eligible_for_bounty").Bool(),
			Type:              item.Get("asset_type").String(),
		}
		scopes = append(scopes, s)
	}
	return scopes
}
