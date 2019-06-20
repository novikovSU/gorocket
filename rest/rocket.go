// Package rest provides a RocketChat rest client.
package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Client AAA
type Client struct {
	Protocol string
	Host     string
	Port     string

	// Use this switch to see all network communication.
	Debug bool

	auth *authInfo
}

type authInfo struct {
	token string
	id    string
}

// SuccessResponse AAA
type SuccessResponse struct {
	Success bool `json:"success"`
}

// The base for the most of the json responses
type statusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// IDOrName helps other functions to choose between RoomID and ChannelName
func IDOrName(id, name string) (_type, value string) {
	if name != "" {
		return "roomName", name
	}
	return "roomId", id
}

// NewClient AAA
func NewClient(host, port string, tls, debug bool) *Client {
	var protocol string

	if tls {
		protocol = "https"
	} else {
		protocol = "http"
	}

	return &Client{Host: host, Port: port, Protocol: protocol, Debug: debug}
}

func (c *Client) getURL() string {
	return fmt.Sprintf("%v://%v:%v", c.Protocol, c.Host, c.Port)
}

func (c *Client) doRequest(request *http.Request, responseBody interface{}) error {

	if c.auth != nil {
		request.Header.Set("X-Auth-Token", c.auth.token)
		request.Header.Set("X-User-Id", c.auth.id)
	}

	if c.Debug {
		log.Println(request)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)

	if c.Debug {
		log.Println(string(bodyBytes))
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("Request error: " + response.Status)
	}

	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, responseBody)
}
