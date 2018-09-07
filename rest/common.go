package rest

import (
	"github.com/killmeplz/gorocket/api"
)

type HistoryOptions struct {
	RoomId    string `url:"roomId"`
	Latest    string `url:"latest,omitempty"`
	Oldest    string `url:"oldest,omitempty"`
	Inclusive bool   `url:"inclusive,omitempty"`
	Count     int64  `url:"count,omitempty"`
	Unreads   bool   `url:"unreads,omitempty"`
}

type History interface {
	History(options *HistoryOptions) ([]api.Message, error)
}
