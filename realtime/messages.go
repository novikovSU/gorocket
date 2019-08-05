package realtime

import (
	"errors"
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

	if rawResponse == nil {
		return nil, errors.New("can't send message")
	}

	return getMessageFromData(rawResponse.(map[string]interface{})), nil
}

// Load a history of a channel
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/load-history/
// RealtimeHistoryRequest AAA

type HistoryOptions struct {
	RoomID     string
	LatestTime *time.Time
	Count      uint
	OldestTime *time.Time
}

type rtHistDate struct {
	datetime int64 `json:"$date"`
}

func (c *Client) LoadHistory(req *HistoryOptions) ([]api.Message, error) {
	var LatestTime *rtHistDate
	var OldestTime *rtHistDate
	if req.LatestTime != nil {
		LatestTime = &rtHistDate{datetime: req.LatestTime.Unix()}
	}
	if req.OldestTime != nil {
		OldestTime = &rtHistDate{datetime: req.OldestTime.Unix()}
	}
	rawResponse, err := c.ddp.Call("loadHistory", req.RoomID, LatestTime, req.Count, OldestTime)
	if err != nil {
		return nil, err
	}

	if rawResponse == nil {
		return nil, errors.New("can't load history")
	}

	resp := gabs.Wrap(rawResponse)
	msgs := resp.S("messages").Children()

	messages := make([]api.Message, 0)

	for _, rawmsg := range msgs {
		msg := getMessageFromDocument(rawmsg)
		if msg != nil {
			messages = append(messages, *msg)
		}
	}

	return messages, nil
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

	messages := make([]api.Message, 0)

	for _, arg := range args {
		msg := getMessageFromDocument(arg)
		if msg != nil {
			messages = append(messages, *msg)
		}
	}

	return messages
}
func getMessageFromDocument(arg *gabs.Container) *api.Message {
	if arg.Path("_id").Data() == nil {
		return nil
	}
	var ts time.Time
	//log.Printf("DEBUG: gabs container: %+v\n", arg)
	tsUnixFloat := arg.Path("ts.$date").Data().(float64)
	tsUnix := int64(tsUnixFloat)
	tsUnixMillis := tsUnix % 1000
	tsUnix = tsUnix / 1000
	ts = time.Unix(tsUnix, tsUnixMillis*1000000)
	return &api.Message{
		ID:        stringOrZero(arg.Path("_id").Data()),
		ChannelID: stringOrZero(arg.Path("rid").Data()),
		Text:      stringOrZero(arg.Path("msg").Data()),
		Timestamp: &ts,
		User: api.User{
			ID:       stringOrZero(arg.Path("u._id").Data()),
			Name:     stringOrZero(arg.Path("u.name").Data()),
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
