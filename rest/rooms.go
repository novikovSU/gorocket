package rest

import (
	"fmt"
	"github.com/killmeplz/gorocket/api"
	"net/http"
)

type Rooms struct {
	client *Client
}

func (c *Client) Rooms() *Rooms {
	return &Rooms{client: c}
}

type RoomResponse struct {
	Update []api.Room `json:"update"`
	Remove []api.Room `json:"remove"`
	SuccessResponse
}

func (r *Rooms) Get() (*RoomResponse, error) {
	var url = fmt.Sprintf("%s/api/v1/rooms.get", r.client.getUrl())
	req, _ := http.NewRequest("GET", url, nil)
	resp := new(RoomResponse)

	err := r.client.doRequest(req, resp)
	return resp, err
}
