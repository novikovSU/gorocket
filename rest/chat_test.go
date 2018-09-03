package rest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRocket_SendAndReceive(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.ListPublicChannels()
	assert.Nil(t, err)

	general := getChannel(rooms, "general")

	err = rocket.Send(general, "Test")
	assert.Nil(t, err)

	messages, err := rocket.ChannelHistory(general, &Page{Count: 10})
	assert.Nil(t, err)

	message := findMessage(messages, testUserName, "Test")
	assert.NotNil(t, message)
}
