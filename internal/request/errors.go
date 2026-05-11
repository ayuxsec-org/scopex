package request

import "errors"

var ErrRateLimited = errors.New("recieved 429 status code")
