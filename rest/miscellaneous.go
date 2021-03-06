package rest

import (
	"net/http"

	"github.com/novikovSU/gorocket/api"
)

type infoResponse struct {
	Info api.Info `json:"info"`
}

// GetServerInfo gets information about the server.
// This function does not need a logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/info
func (c *Client) GetServerInfo() (*api.Info, error) {
	request, _ := http.NewRequest("GET", c.getURL()+"/api/v1/info", nil)

	response := new(infoResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Info, nil
}
