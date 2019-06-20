package rest

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/novikovSU/gorocket/api"
)

type groupsResponse struct {
	Success bool        `json:"success"`
	Groups  []api.Group `json:"groups"`
}

type groupResponse struct {
	Success bool      `json:"success"`
	Group   api.Group `json:"group"`
}

// Groups AAA
type Groups struct {
	client *Client
}

// Groups AAA
func (client *Client) Groups() *Groups {
	return &Groups{client: client}
}

// ListGroups AAA
func (g *Groups) ListGroups() ([]api.Group, error) {
	request, _ := http.NewRequest(http.MethodGet, g.client.getURL()+"/api/v1/groups.list", nil)
	response := new(groupsResponse)

	if err := g.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Groups, nil
}

// GroupsInfoOptions AAA
type GroupsInfoOptions struct {
	RoomID   string `url:"roomId,omitempty"`
	RoomName string `url:"roomName,omitempty"`
}

// Info AAA
func (g *Groups) Info(opts *GroupsInfoOptions) (*api.Group, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	var url = fmt.Sprintf("%s/api/v1/groups.info?%s", g.client.getURL(), vals.Encode())
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response := new(groupResponse)

	err = g.client.doRequest(request, response)
	fmt.Println(response)

	return &response.Group, err
}

// History AAA
func (g *Groups) History(opts *HistoryOptions) ([]api.Message, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/groups.history?%s", g.client.getURL(), vals.Encode())

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response := new(MessagesResponse)

	if err := g.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}
