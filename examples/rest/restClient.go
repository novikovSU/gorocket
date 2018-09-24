package main

import (
	"github.com/killmeplz/gorocket/api"
	"github.com/killmeplz/gorocket/rest"
	"log"
)

func main() {
	// Create a gorocket rest client
	c := rest.NewClient("127.0.0.1", "3000", false, false)
	// close client
	//defer c.Close()

	// No login needed to get the server versions
	info, _ := c.Misc().GetServerInfo()
	log.Println("Server version: ", info.Info.Version)

	// Login an existing user
	if err := c.Login(api.UserCredentials{Email: "test@mail.com", Password: "test"}); err != nil {
		log.Fatal("Error while login: ", err)
	}

	// Get all visible channels
	channels, _ := c.Channel().List()
	log.Println("All visible channels: ", channels)

	// Join the general channel
	general := getChannelById(channels, "GENERAL")

	// Get all joined channels
	joined, _ := c.Channel().ListJoined()
	log.Println("We are in the following channels: ", joined)

	// Send a message
	c.Chat().Post(&rest.ChatPostOptions{RoomId: general.Id, Text: "Test"})

	// Get the last messages from the general channel
	messages, _ := c.Channel().History(&rest.HistoryOptions{RoomId: general.Id})
	log.Println("Last messages: ", messages)

	// Leave the general channel
	c.Channel().Leave(&rest.ChannelLeaveOptions{RoomId: general.Id})

	// Logout the user
	c.Logout()
}

func getChannelById(channels []api.Channel, id string) *api.Channel {
	for _, c := range channels {
		if c.Id == id {
			return &c
		}
	}

	return nil
}
