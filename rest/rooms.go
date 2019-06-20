package rest

import (
	"fmt"
	"net/http"

	"github.com/novikovSU/gorocket/api"
)

// Rooms AAA
type Rooms struct {
	client *Client
}

// Rooms AAA
func (c *Client) Rooms() *Rooms {
	return &Rooms{client: c}
}

// RoomResponse AAA
type RoomResponse struct {
	Update []api.Room `json:"update"`
	Remove []api.Room `json:"remove"`
	SuccessResponse
}

// Get AAA
func (r *Rooms) Get() ([]api.Room, error) {
	var url = fmt.Sprintf("%s/api/v1/rooms.get", r.client.getURL())
	request, _ := http.NewRequest("GET", url, nil)
	response := new(RoomResponse)

	err := r.client.doRequest(request, response)
	return response.Update, err
}
