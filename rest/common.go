package rest

import (
	"github.com/novikovSU/gorocket/api"
)

// HistoryOptions AAA
type HistoryOptions struct {
	RoomID    string `url:"roomId"`
	Latest    string `url:"latest,omitempty"`
	Oldest    string `url:"oldest,omitempty"`
	Inclusive bool   `url:"inclusive,omitempty"`
	Count     int64  `url:"count,omitempty"`
	Unreads   bool   `url:"unreads,omitempty"`
}

// History AAA
type History interface {
	History(options *HistoryOptions) ([]api.Message, error)
}
