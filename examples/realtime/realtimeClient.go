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
	c, _ := realtime.NewClient("ws", "rocket.example.com", "3000", true)
	// close the client
	defer c.Close()

	// Login an existing user
	c.Login(&api.UserCredentials{Email: "tester@example.com", Name: "tester", Password: "qwe123"})

	// Subscribe to the general channel
	allMessages := api.Channel{ID: "__my_messages__"}
	messageChannel, _ := c.SubscribeToMessageStream(&allMessages)

	// Send some messages
	sendChan := api.Channel{ID: "3D3r4y9Yn8AEwqejo"}
	c.SendMessage(&sendChan, "This")
	c.SendMessage(&sendChan, "is")
	c.SendMessage(&sendChan, "sparta!")

	go func() {
		for {
			msg := <-messageChannel
			fmt.Printf("MESSAGE: %+v\n", msg)
		}
	}()

	for {
	}
}
