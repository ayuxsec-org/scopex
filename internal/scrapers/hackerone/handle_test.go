package hackerone

import "testing"

func TestGetAllHandles(t *testing.T) {
	handles, err := scraper.GetAllHandles()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if len(handles) == 0 {
		t.Fatalf("no handles found")
	}
	for _, handle := range handles {
		t.Log(handle)
	}
}
