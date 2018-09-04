package rest

import (
	"bytes"
	"fmt"
	"github.com/killmeplz/gorocket/api"
	"net/http"
	"encoding/json"
	"github.com/google/go-querystring/query"
)

type ChannelsResponse struct {
	Success  bool          `json:"success"`
	Channels []api.Channel `json:"channels"`
}

type ChannelResponse struct {
	Success bool        `json:"success"`
	Channel api.Channel `json:"channel"`
}

type Channel struct {
	client *Client
}

func (c *Client) Channel() *Channel {
	return &Channel{ client : c }
}

type ChannelOptions struct {
	RoomId string `json:"roomId,omitempty"`
	RoomName string `json:"roomName,omitempty"`
}

// Returns all channels that can be seen by the logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list
func (c *Client) ListPublicChannels() ([]api.Channel, error) {
	request, _ := http.NewRequest("GET", c.getUrl()+"/api/v1/channels.list", nil)
	response := new(ChannelsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

// Returns all channels that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list-joined
func (c *Client) ListJoinedChannels() ([]api.Channel, error) {
	request, _ := http.NewRequest("GET", c.getUrl()+"/api/v1/channels.list.joined", nil)
	response := new(ChannelsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}


type ChannelLeaveOptions struct {
	RoomId string `json:"roomId,omitempty"`
}
// Leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (c *Channel) Leave(opts *ChannelLeaveOptions) (*ChannelResponse , error) {
	data, err := json.Marshal(&opts)
	if err != nil {
		return nil,err
	}
	request, _ := http.NewRequest("POST", c.client.getUrl()+"/api/v1/channels.leave", bytes.NewBuffer(data))
	var resp ChannelResponse
	err = c.client.doRequest(request, resp)
	return &resp, err
}



// Get information about a channel. That might be useful to update the usernames.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/info
func (c *Channel) Info(opts *ChannelOptions) (*api.Channel, error) {
	room , nameOrId := IdOrName(opts.RoomId,opts.RoomName)
	var url = fmt.Sprintf("%s/api/v1/channels.info?%s=%s", c.client.getUrl(), room, nameOrId )
	request, _ := http.NewRequest("GET", url, nil)
	response := new(ChannelResponse)

	if err := c.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Channel, nil
}

type ChannelHistoryOptions struct {
	RoomId string `url:"roomId"`
	Latest string `url:"latest,omitempty"`
	Oldest string `url:"oldest,omitempty"`
	Inclusive bool `url:"inclusive,omitempty"`
	Count int64 `url:"count,omitempty"`
	Unreads bool `url:"unreads,omitempty"`

}

// Get messages from a channel. The channel id has to be not nil. Optionally a
// count can be specified to limit the size of the returned messages.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/history
func (c *Channel) History(opts *ChannelHistoryOptions) ([]api.Message, error) {
	vals , err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/channels.history?%s", c.client.getUrl(), vals.Encode())

	request, _ := http.NewRequest("GET", url, nil)
	response := new(MessagesResponse)

	if err := c.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}
