package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/novikovSU/gorocket/api"
)

// MessagesResponse AAA
type MessagesResponse struct {
	statusResponse
	ChannelName string        `json:"channel"`
	Messages    []api.Message `json:"messages"`
}

// MessageResponse AAA
type MessageResponse struct {
	statusResponse
	ChannelName string      `json:"channel"`
	Message     api.Message `json:"message"`
	Success     bool        `json:"success"`
	TimeStamp   int64       `json:"ts"`
}

// Page AAA
type Page struct {
	Count int
}

// Chat AAA
type Chat struct {
	client *Client
}

// Chat AAA
func (c *Client) Chat() *Chat {
	return &Chat{c}
}

// ChatPostOptions PostMessage Payload. The name of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
type ChatPostOptions struct {
	RoomID      string           `json:"roomId,omitempty"`
	Channel     string           `json:"channel,omitempty"`
	Text        string           `json:"text,omitempty"`
	Alias       string           `json:"alias,omitempty"`
	Emoji       string           `json:"emoji,omitempty"`
	Avatar      string           `json:"avatar,omitempty"`
	Attachments []api.Attachment `json:"attachments,omitempty"`
}

// Post sends a message to a channel. The name of the channel has to be not nil.
// The message will be html escaped.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
func (c *Chat) Post(opts *ChatPostOptions) (*MessageResponse, error) {
	data, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest("POST", c.client.getURL()+"/api/v1/chat.postMessage", bytes.NewBuffer(data))

	response := new(MessageResponse)
	err = c.client.doRequest(request, response)

	return response, err
}

// ReceiptsResponse AAA
type ReceiptsResponse struct {
	Receipts []api.ReadReceipt `json:"receipts"`
	Success  bool              `json:"success"`
}

// GetMessageReadReceipts AAA
func (c *Chat) GetMessageReadReceipts(msgID string) ([]api.ReadReceipt, error) {
	req, _ := http.NewRequest(
		http.MethodGet,
		c.client.getURL()+"/api/v1/chat.getMessageReadReceipts?messageId="+msgID,
		nil)
	resp := new(ReceiptsResponse)
	err := c.client.doRequest(req, resp)
	return resp.Receipts, err
}

// ChatUpdateOptions UpdateMessage Payload
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/update/
type ChatUpdateOptions struct {
	RoomID string `json:"roomId"`
	MsgID  string `json:"msgId"`
	Text   string `json:"text"`
}

// Update the text of the chat message.. The name of the channel has to be not nil.
// The message will be html escaped.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/update/
func (c *Chat) Update(opts *ChatUpdateOptions) (*MessageResponse, error) {
	data, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest("POST", c.client.getURL()+"/api/v1/chat.update", bytes.NewBuffer(data))

	response := new(MessageResponse)
	err = c.client.doRequest(request, response)

	return response, err
}
