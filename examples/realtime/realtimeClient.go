package main

import (
	"fmt"

	"github.com/novikovSU/gorocket/api"
	"github.com/novikovSU/gorocket/realtime"
)

func main() {
	// Create new realtime client
	// Use "ws" as a proto for unencrypted connection
	// and "wss" for encrypted
	c, _ := realtime.NewClient("ws", "127.0.0.1", "3000", false)
	// close the client
	defer c.Close()

	// Login an existing user
	c.Login(&api.UserCredentials{Email: "reatimeTest@mail.com", Name: "realtime", Password: "realtime"})

	// Subscribe to the general channel
	general := api.Channel{ID: "GENERAL"}
	messageChannel, _ := c.SubscribeToMessageStream(&general)

	// Send some messages
	c.SendMessage(&general, "This")
	c.SendMessage(&general, "is")
	c.SendMessage(&general, "sparta!")

	// Get messages from channel
	fmt.Println(<-messageChannel)
	fmt.Println(<-messageChannel)
	fmt.Println(<-messageChannel)

}
