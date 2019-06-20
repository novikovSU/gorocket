package realtime

import (
	"github.com/novikovSU/gorocket/api"
	"github.com/novikovSU/gorocket/common_testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	client *Client
)

func getLoggedInClient(t *testing.T) *Client {

	if client == nil {
		c, err := NewClient(common_testing.Host, common_testing.Port, true)
		assert.Nil(t, err, "Couldn't create realtime client")

		err = c.RegisterUser(&api.UserCredentials{
			Email:    common_testing.GetRandomEmail(),
			Name:     common_testing.GetRandomString(),
			Password: common_testing.GetRandomString()})
		assert.Nil(t, err, "Couldn't register user")

		client = c
	}

	return client
}
