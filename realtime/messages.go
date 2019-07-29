package realtime

import (
	"fmt"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/gopackage/ddp"
	"github.com/novikovSU/gorocket/api"
)

const (
	// RocketChat doesn't send the `added` event for new messages by default, only `changed`.
	send_added_event    = true
	default_buffer_size = 100
)

// Send a message to a channel
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/send-message/
func (c *Client) SendMessage(channel *api.Channel, text string) (*api.Message, error) {
	m := api.Message{
		ID:        c.newRandomId(),
		ChannelID: channel.ID,
		Text:      text}

	rawResponse, err := c.ddp.Call("sendMessage", m)

	if err != nil {
		return nil, err
	}

	return getMessageFromData(rawResponse.(map[string]interface{})), nil
}

// Subscribes to the message updates of a channel
// Returns a buffered channel
//
// https://rocket.chat/docs/developer-guides/realtime-api/subscriptions/stream-room-messages/
func (c *Client) SubscribeToMessageStream(channel *api.Channel) (chan api.Message, error) {

	if err := c.ddp.Sub("stream-room-messages", channel.ID, send_added_event); err != nil {
		return nil, err
	}

	msgChannel := make(chan api.Message, default_buffer_size)
	c.ddp.CollectionByName("stream-room-messages").AddUpdateListener(messageExtractor{msgChannel, "update"})

	return msgChannel, nil
}

func getMessageFromData(data interface{}) *api.Message {
	document := gabs.Wrap(data)
	return getMessageFromDocument(document)
}

func getMessagesFromUpdateEvent(update ddp.Update) []api.Message {
	document := gabs.Wrap(update["args"])
	args := document.Children()

	messages := make([]api.Message, len(args))

	for i, arg := range args {
		messages[i] = *getMessageFromDocument(arg)
	}

	return messages
}
func getMessageFromDocument(arg *gabs.Container) *api.Message {
	tsUnix := int64(arg.Path("ts.$date").Data().(float64))
	tsUnixMillis := tsUnix % 1000
	tsUnix = tsUnix / 1000
	ts := time.Unix(tsUnix, tsUnixMillis*1000000)
	return &api.Message{
		ID:        stringOrZero(arg.Path("_id").Data()),
		ChannelID: stringOrZero(arg.Path("rid").Data()),
		Text:      stringOrZero(arg.Path("msg").Data()),
		Timestamp: &ts,
		User: api.User{
			ID:       stringOrZero(arg.Path("u._id").Data()),
			UserName: stringOrZero(arg.Path("u.username").Data())}}
}

func stringOrZero(i interface{}) string {
	if i == nil {
		return ""
	}

	switch i.(type) {
	case string:
		return i.(string)
	case float64:
		return fmt.Sprintf("%f", i.(float64))
	default:
		return ""
	}
}

type messageExtractor struct {
	messageChannel chan api.Message
	operation      string
}

func (u messageExtractor) CollectionUpdate(collection, operation, id string, doc ddp.Update) {
	if operation == u.operation {
		msgs := getMessagesFromUpdateEvent(doc)
		for _, m := range msgs {
			u.messageChannel <- m
		}
	}
}
