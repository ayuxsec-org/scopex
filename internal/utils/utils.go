package utils

import "github.com/ayuxsec-org/log"

func Must[T any](v T, err error) T {
	if err != nil {
		log.Fatalf("'utils.Must' error: %v", err)
	}
	return v
}
