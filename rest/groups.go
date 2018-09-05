package rest

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/killmeplz/gorocket/api"
	"net/http"
)

type groupsResponse struct {
	Success bool        `json:"success"`
	Groups  []api.Group `json:"groups"`
}

type groupResponse struct {
	Success bool      `json:"success"`
	Group   api.Group `json:"group"`
}

type Groups struct {
	client *Client
}

func (client *Client) Groups() *Groups {
	return &Groups{client: client}
}

func (g *Groups) ListGroups() ([]api.Group, error) {
	request, _ := http.NewRequest(http.MethodGet, g.client.getUrl()+"/api/v1/groups.list", nil)
	response := new(groupsResponse)

	if err := g.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Groups, nil
}

type GroupsInfoOptions struct {
	RoomId   string `url:"roomId,omitempty"`
	RoomName string `url:"roomName,omitempty"`
}

func (g *Groups) Info(opts *GroupsInfoOptions) (*api.Group, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	var url = fmt.Sprintf("%s/api/v1/groups.info?%s", g.client.getUrl(), vals.Encode())
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response := new(groupResponse)

	err = g.client.doRequest(request, response)
	fmt.Println(response)

	return &response.Group, err
}

type GroupsHistoryOptions struct {
	RoomId    string `url:"roomId"`
	Latest    string `url:"latest,omitempty"`
	Oldest    string `url:"oldest,omitempty"`
	Inclusive bool   `url:"inclusive,omitempty"`
	Count     int64  `url:"count,omitempty"`
	Unreads   bool   `url:"unreads,omitempty"`
}

func (g *Groups) History(opts *GroupsHistoryOptions) ([]api.Message, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/groups.history?%s", g.client.getUrl(), vals.Encode())

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response := new(MessagesResponse)

	if err := g.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}
