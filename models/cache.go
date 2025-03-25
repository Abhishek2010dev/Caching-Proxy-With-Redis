package models

import (
	"time"
)

type CachedEntry struct {
	StatusCode   int                 `json:"statusCode"`
	Header       map[string][]string `json:"header"`
	ResponseBody []byte              `json:"responseBody"`
	Created      time.Time           `json:"created"`
}
