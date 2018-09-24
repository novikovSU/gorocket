package rest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRocket_SendAndReceive(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.Channel().List()
	assert.Nil(t, err)

	general := getChannel(rooms, "general")

	_, err = rocket.Chat().Post(&ChatPostOptions{Channel: "general", Text: "Test"})
	assert.Nil(t, err)
	messages, err := rocket.Channel().History(&HistoryOptions{RoomId: general.Id, Count: 10})
	assert.Nil(t, err)

	message := findMessage(messages, testUserName, "Test")
	assert.NotNil(t, message)
}
