package rest

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"

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

// RoomsUploadOptions AAA
type RoomsUploadOptions struct {
	//	RoomID      string `json:"rid"`
	File        string `json:"file"`
	Message     string `json:"msg,omitempty"`
	Description string `json:"description,omitempty"`
}

// RoomsUploadResponse AAA
type RoomsUploadResponse struct {
	SuccessResponse
}

// Upload sends a message with a file to a room. The name of the room has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/rooms/upload/
func (r *Rooms) Upload(roomID string, opts *RoomsUploadOptions) (*RoomsUploadResponse, error) {
	fileDir, _ := os.Getwd()
	fileName := opts.File
	filePath := path.Join(fileDir, fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("msg", opts.Message)
	_ = writer.WriteField("description", opts.Description)
	writer.Close()

	request, _ := http.NewRequest("POST", r.client.getURL()+"/api/v1/rooms.upload/"+roomID, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	response := new(RoomsUploadResponse)
	err := r.client.doRequest(request, response)

	return response, err
}

// Get AAA
func (r *Rooms) Get() ([]api.Room, error) {
	var url = fmt.Sprintf("%s/api/v1/rooms.get", r.client.getURL())
	request, _ := http.NewRequest("GET", url, nil)
	response := new(RoomResponse)

	err := r.client.doRequest(request, response)
	return response.Update, err
}
