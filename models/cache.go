package models

import (
	"time"
)

type CachedEntry struct {
	ResponseBody []byte    `json:"responseBody"`
	Created      time.Time `json:"created"`
}
