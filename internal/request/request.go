package request

import (
	"fmt"
	"net/http"

	"github.com/ayuxsec-org/scopex/internal/utils"
)

type Request struct {
	UserName string
	Password string
	Client   *http.Client
}

// caller must close the body
func (r *Request) Get(url string) (*http.Response, error) {
	req := utils.Must(http.NewRequest(http.MethodGet, url, nil))
	req.SetBasicAuth(r.UserName, r.Password)
	resp, err := r.Client.Do(req)
	if err != nil {
		return &http.Response{}, fmt.Errorf("'r.Client.Do' error: %v", err)
	}
	if resp.StatusCode == 429 {
		return &http.Response{}, ErrRateLimited
	}
	return resp, nil
}
