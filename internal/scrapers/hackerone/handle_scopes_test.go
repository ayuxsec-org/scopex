package hackerone

import (
	"testing"
)

func TestGetAllHandlesAndTheirScopes(t *testing.T) {
	scopes, err := scraper.GetAllHandlesAndTheirScopes()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(scopes)
}

func TestGetHandleScopes(t *testing.T) {
	scopes, err := scraper.GetHandleScopes("gitlab")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	t.Log(scopes)
}
