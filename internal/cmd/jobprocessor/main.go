package main

import (
	"log"

	"go.kimpton.io/url-checker/internal/bootstrap"
	"go.kimpton.io/url-checker/internal/config"
	"go.kimpton.io/url-checker/internal/domain"
	"go.kimpton.io/url-checker/internal/tasks"
	"go.kimpton.io/url-checker/internal/workers"
)

func main() {

	log.Printf("Starting job processor...\n")

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// Get a message queue provider
	mq, err := bootstrap.GetMessageQueueProvider(conf)
	if err != nil {
		panic(err)
	}

	// Build a worker to consume messages from the queue
	w := workers.NewMQWorker(mq)
	task := &tasks.HTTPChecker{}
	go w.Start(domain.HTTPSCheck, task)

	// Keep the server running
	select {}

}
