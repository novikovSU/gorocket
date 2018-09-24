package rest

import (
	"github.com/killmeplz/gorocket/api"
	"net/http"
)

type Misc struct {
	client *Client
}

func (c *Client) Misc() *Misc {
	return &Misc{client: c}
}

type ServerInfoResponse struct {
	Info api.Info `json:"info"`
	SuccessResponse
}

// Get information about the server.
// This function does not need a logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/info
func (m *Misc) GetServerInfo() (*ServerInfoResponse, error) {
	req, _ := http.NewRequest("GET", m.client.getUrl()+"/api/v1/info", nil)

	resp := new(ServerInfoResponse)

	if err := m.client.doRequest(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
