package rest

import (
	"fmt"
	"net/http"
	"github.com/killmeplz/gorocket/api"
)

type Rooms struct {
	client *Client
}

func (c *Client) Rooms() *Rooms {
	return &Rooms{client:c}
}



type RoomResponse struct {
	Update []api.Room `json:"update"`
	Remove []api.Room `json:"remove"`
	SuccessResponse
}

func (r *Rooms) Get() ([]api.Room,error) {
	var url = fmt.Sprintf("%s/api/v1/rooms.get", r.client.getUrl())
	request, _ := http.NewRequest("GET", url, nil)
	response := new(RoomResponse)

	err := r.client.doRequest(request, response)
	return response.Update, err
}