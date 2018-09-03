package rest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRocket_GetPublicChannels(t *testing.T) {
	rocket := getDefaultClient(t)

	channels, err := rocket.ListPublicChannels()
	assert.Nil(t, err)

	assert.Len(t, channels, 1)
	assert.Equal(t, "general", channels[0].Name)
}

func TestRocket_GetJoinedChannels(t *testing.T) {
	rocket := getDefaultClient(t)

	channels, err := rocket.ListPublicChannels()
	assert.Nil(t, err)

	channels, err = rocket.ListJoinedChannels()
	assert.Nil(t, err)

	assert.Len(t, channels, 1)
	assert.Equal(t, "general", channels[0].Name)
}

func TestRocket_LeaveChannel(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.ListPublicChannels()
	assert.Nil(t, err)

	general := getChannel(rooms, "general")

	err = rocket.LeaveChannel(general)
	assert.Nil(t, err)
}

func TestRocket_GetChannelInfo(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.ListPublicChannels()
	assert.Nil(t, err)

	general := getChannel(rooms, "general")

	updatedChannelInfo, err := rocket.GetChannelInfo(general)
	assert.Nil(t, err)
	assert.NotNil(t, updatedChannelInfo)

	assert.Equal(t, general.Id, updatedChannelInfo.Id)
	assert.NotEmpty(t, updatedChannelInfo.Name)
	assert.NotEmpty(t, updatedChannelInfo.T)
	assert.NotEmpty(t, updatedChannelInfo.UpdatedAt)
	assert.NotEmpty(t, updatedChannelInfo.Timestamp)
	assert.NotZero(t, len(updatedChannelInfo.UserNames))
}
