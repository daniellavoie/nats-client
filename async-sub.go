package main

import (
	"log"
	"runtime"

	"github.com/nats-io/go-nats"
	"os"
)

func main() {
	// Create server connection
	url := os.Args[2]
	if url == "" {
		url = nats.DefaultURL
	}
	natsConnection, _ := nats.Connect(url)
	log.Println("Connected to " + url)

	// Subscribe to subject
	subject := os.Args[1]
	if subject == "" {
		log.Printf("Error : No subject specified.")

		os.Exit(1)
	}

	log.Printf("Subscribing to subject '%s'\n", subject)
	natsConnection.Subscribe("subject", func(msg *nats.Msg) {

		// Handle the message
		log.Printf("Received message '%s\n", string(msg.Data)+"'")
	})

	// Keep the connection alive
	runtime.Goexit()
}
