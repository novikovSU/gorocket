package rest

import (
	"bytes"
	"fmt"
	"github.com/killmeplz/gorocket/api"
	"html"
	"net/http"
)

type messagesResponse struct {
	statusResponse
	ChannelName string        `json:"channel"`
	Messages    []api.Message `json:"messages"`
}

type messageResponse struct {
	statusResponse
	ChannelName string      `json:"channel"`
	Message     api.Message `json:"message"`
}

type Page struct {
	Count int
}

type Chat struct {
	client *Client
}

func (c *Client) Chat() Chat {
	return Chat{c}
}

// Sends a message to a channel. The name of the channel has to be not nil.
// The message will be html escaped.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
func (c Chat) Send(channel *api.Channel, msg string) error {
	body := fmt.Sprintf(`{ "channel": "%s", "text": "%s"}`, channel.Name, html.EscapeString(msg))
	request, _ := http.NewRequest("POST", c.client.getUrl()+"/api/v1/chat.postMessage", bytes.NewBufferString(body))

	response := new(messageResponse)

	return c.client.doRequest(request, response)
}

// Updates the text of the chat message.. The name of the channel has to be not nil.
// The message will be html escaped.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/update/
func (c Chat)  Update(channel *api.Channel, msgid, msg string) error {
	body := fmt.Sprintf(`{ "roomId": "%s", "msgId": "%s", "text": "%s"}`, channel.Id, msgid, html.EscapeString(msg))
	request, _ := http.NewRequest("POST", c.client.getUrl()+"/api/v1/chat.update", bytes.NewBufferString(body))

	response := new(messageResponse)

	return c.client.doRequest(request, response)
}