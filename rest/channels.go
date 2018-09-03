package rest

import (
	"bytes"
	"fmt"
	"github.com/killmeplz/gorocket/api"
	"net/http"
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

// Leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (c *Client) LeaveChannel(channel *api.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, channel.Id)
	request, _ := http.NewRequest("POST", c.getUrl()+"/api/v1/channels.leave", bytes.NewBufferString(body))
	return c.doRequest(request, new(statusResponse))
}

type ChannelGetInfoOptions struct {
	RoomId string `json:"roomId,omitempty"`
	RoomName string `json:"roomName,omitempty"`
}

// Get information about a channel. That might be useful to update the usernames.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/info
func (c *Channel) Info(opts *ChannelGetInfoOptions) (*api.Channel, error) {
	room , nameOrId := IdOrName(opts.RoomId,opts.RoomName)
	var url = fmt.Sprintf("%s/api/v1/channels.info?%s=%s", c.client.getUrl(), room, nameOrId )
	request, _ := http.NewRequest("GET", url, nil)
	response := new(ChannelResponse)

	if err := c.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Channel, nil
}

// Get messages from a channel. The channel id has to be not nil. Optionally a
// count can be specified to limit the size of the returned messages.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/history
func (c *Client) ChannelHistory(channel *api.Channel, page *Page) ([]api.Message, error) {
	u := fmt.Sprintf("%s/api/v1/channels.history?roomId=%s", c.getUrl(), channel.Id)

	if page != nil {
		u = fmt.Sprintf("%s&count=%d", u, page.Count)
	}

	request, _ := http.NewRequest("GET", u, nil)
	response := new(MessagesResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}
