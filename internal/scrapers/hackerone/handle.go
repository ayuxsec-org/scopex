package hackerone

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/ayuxsec-org/log"
	"github.com/ayuxsec-org/scopex/internal/utils"
	"github.com/tidwall/gjson"
)

const (
	// MaxHandleReachedIdentifier is the response returned when no more handles are available
	MaxHandleReachedIdentifier string = `{"data":[],"links":{}}`
)

type Handle struct {
	Handle string `json:"handle"`
	IsBBP  bool   `json:"is_bbp"`
}

// GetAllHandles fetches the program handles from HackerOne
// It paginates through all the pages to fetch all the handles until it encounters an empty response
func (s *Scraper) GetAllHandles() (handles []Handle, err error) {
	var counter int
	for {
		resp, err := s.r.Get(ProgramHandlesAPIUrl + strconv.Itoa(counter))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", ErrSendingRequest, err)
		}
		defer resp.Body.Close()

		respBody := string(utils.Must(io.ReadAll(resp.Body)))

		if strings.Contains(respBody, MaxHandleReachedIdentifier) {
			break
		}
		handles = append(handles, extractHandles(respBody)...)
		counter += 1
	}
	log.Infof("Scraped %d handles", len(handles))
	return
}

// extractHandles extracts the handles from the JSON response
func extractHandles(jsonStr string) (handles []Handle) {
	results := gjson.Get(jsonStr, "data.#.attributes")
	for _, item := range results.Array() {
		h := Handle{
			Handle: item.Get("handle").String(),
			IsBBP:  item.Get("offers_bounties").Bool(),
		}
		if !slices.Contains(handles, h) {
			handles = append(handles, h)
		}
	}
	return
}
