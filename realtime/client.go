// Provides access to the Meteor ddp realtime API of Rocket.Chat
package realtime

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gopackage/ddp"
)

type Client struct {
	ddp *ddp.Client
}

// Creates a new instance and connects to the websocket.
func NewClient(proto, host, port string, debug bool) (*Client, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	c := new(Client)
	c.ddp = ddp.NewClient(fmt.Sprintf("%v://%v:%v/websocket", proto, host, port), "http://"+host)

	if debug {
		c.ddp.SetSocketLogActive(true)
	}

	if err := c.ddp.Connect(); err != nil {
		return nil, err
	}

	return c, nil
}

// Closes the ddp session
func (c *Client) Close() {
	c.ddp.Close()
}

// Some of the rocketchat objects need unique IDs specified by the client
func (c *Client) newRandomId() string {
	return fmt.Sprintf("%f", rand.Float64())
}
