package hackerone

import "errors"

// ErrSendingRequest is returned when there is an error while sending the request to the HackerOne API
var ErrSendingRequest = errors.New("error while sending request to hackerone API")

// ErrGettingProgramHandles is returned when there is an error while getting program handles
var ErrGettingProgramHandles = errors.New("error while enumerating program handles")

// ErrGettingProgramScopes is returned when there is an error while getting program scopes/assets
var ErrGettingProgramScopes = errors.New("error getting program scopes/assets")
