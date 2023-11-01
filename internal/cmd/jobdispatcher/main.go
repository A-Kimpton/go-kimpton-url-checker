package main

import (
	"encoding/json"
	"log"

	"go.kimpton.io/url-checker/internal/bootstrap"
	"go.kimpton.io/url-checker/internal/config"
	"go.kimpton.io/url-checker/internal/domain"
	"go.kimpton.io/url-checker/internal/domain/messages"
)

func main() {

	log.Print("Starting job dispatcher...\n")

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// Get a message queue provider
	mq, err := bootstrap.GetMessageQueueProvider(conf)
	if err != nil {
		panic(err)
	}

	// Build a message for the queue
	data, err := json.Marshal(messages.HTTPSRequest{
		Url: "https://www.google.com",
	})
	if err != nil {
		panic(err)
	}

	// Publish the message
	err = mq.Publish(domain.HTTPSCheck, data)
	if err != nil {
		panic(err)
	}
	log.Println("Published message")

}
