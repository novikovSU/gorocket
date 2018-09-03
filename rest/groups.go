package rest

import (
	"github.com/killmeplz/gorocket/api"
	"net/http"
)

type groupsResponse struct {
	Success bool        `json:"success"`
	Groups  []api.Group `json:"channels"`
}

type groupResponse struct {
	Success bool      `json:"success"`
	Group   api.Group `json:"channel"`
}

func (c *Client) ListGroups() ([]api.Group, error) {
	request, _ := http.NewRequest("GET", c.getUrl()+"/api/v1/groups.list", nil)
	response := new(groupsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Groups, nil
}
