package rest

import (
	"github.com/novikovSU/gorocket/api"
	"github.com/novikovSU/gorocket/common_testing"
	"github.com/novikovSU/gorocket/realtime"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testUserName  string
	testUserEmail string
	testPassword  = "test"
	rocketClient  *Client
)

func getDefaultClient(t *testing.T) *Client {

	if rocketClient == nil {
		testUserEmail = common_testing.GetRandomEmail()
		testUserName = common_testing.GetRandomString()
		rocketClient = getAuthenticatedClient(t, testUserName, testUserEmail, testPassword)
	}

	return rocketClient
}

func getAuthenticatedClient(t *testing.T, name, email, password string) *Client {
	client := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}
	credentials := api.UserCredentials{Name: name, Email: email, Password: password}

	rtClient, err := realtime.NewClient(common_testing.Host, common_testing.Port, true)
	assert.Nil(t, err)
	regErr := rtClient.RegisterUser(&credentials)
	assert.Nil(t, regErr)

	loginErr := client.Login(credentials)
	assert.Nil(t, loginErr)

	return &client
}

func findMessage(messages []api.Message, user string, msg string) *api.Message {
	for _, m := range messages {
		if m.User.UserName == user && m.Text == msg {
			return &m
		}
	}

	return nil
}

func getChannel(channels []api.Channel, name string) *api.Channel {
	for _, r := range channels {
		if r.Name == name {
			return &r
		}
	}

	return nil
}
