package models

import (
	"net/http"
	"time"
)

type CachedEntry struct {
	Response     *http.Response `json:"response"`
	ResponseBody []byte         `json:"responseBody"`
	Created      time.Time      `json:"created"`
}
