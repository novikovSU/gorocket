package rest

import (
	"bytes"
	"encoding/json"
	"github.com/killmeplz/gorocket/api"
	"net/http"
)

type MessagesResponse struct {
	statusResponse
	ChannelName string        `json:"channel"`
	Messages    []api.Message `json:"messages"`
}

type MessageResponse struct {
	statusResponse
	ChannelName string      `json:"channel"`
	Message     api.Message `json:"message"`
	Success     bool        `json:"success"`
	TimeStamp   int64       `json:"ts"`
}

type Page struct {
	Count int
}

type Chat struct {
	client *Client
}

func (c *Client) Chat() *Chat {
	return &Chat{c}
}

// PostMessage Payload. The name of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
type ChatPostOptions struct {
	RoomId      string           `json:"roomId,omitempty"`
	Channel     string           `json:"channel,omitempty"`
	Text        string           `json:"text,omitempty"`
	Alias       string           `json:"alias,omitempty"`
	Emoji       string           `json:"emoji,omitempty"`
	Avatar      string           `json:"avatar,omitempty"`
	Attachments []api.Attachment `json:"attachments,omitempty"`
}

// Sends a message to a channel. The name of the channel has to be not nil.
// The message will be html escaped.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
func (c *Chat) Post(opts *ChatPostOptions) (*MessageResponse, error) {
	data, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest("POST", c.client.getUrl()+"/api/v1/chat.postMessage", bytes.NewBuffer(data))

	response := new(MessageResponse)
	err = c.client.doRequest(request, response)

	return response, err
}

type ReceiptsResponse struct {
	Receipts []api.ReadReceipt `json:"receipts"`
	Success  bool              `json:"success"`
}

func (c *Chat) GetMessageReadReceipts(msgId string) ([]api.ReadReceipt, error) {
	req, _ := http.NewRequest(
		http.MethodGet,
		c.client.getUrl()+"/api/v1/chat.getMessageReadReceipts?messageId="+msgId,
		nil)
	resp := new(ReceiptsResponse)
	err := c.client.doRequest(req, resp)
	return resp.Receipts, err
}

// UpdateMessage Payload
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/update/
type ChatUpdateOptions struct {
	RoomId string `json:"roomId"`
	MsgId  string `json:"msgId"`
	Text   string `json:"text"`
}

// Updates the text of the chat message.. The name of the channel has to be not nil.
// The message will be html escaped.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/update/
func (c *Chat) Update(opts *ChatUpdateOptions) (*MessageResponse, error) {
	data, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest("POST", c.client.getUrl()+"/api/v1/chat.update", bytes.NewBuffer(data))

	response := new(MessageResponse)
	err = c.client.doRequest(request, response)

	return response, err
}
