package workers

import (
	"log"

	"kimpton.io/url-checker/internal/domain"
)

type MQWorker struct {
	mq domain.MessageQueue
}

func NewMQWorker(mq domain.MessageQueue) *MQWorker {
	return &MQWorker{mq}
}

func (w *MQWorker) Start(queueName domain.QueueName, task domain.Task) {
	messages, err := w.mq.Consume(queueName)
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}

	for msg := range messages {
		if _, err := task.Execute(msg); err != nil {
			log.Println("Failed to execute task:", err)
		}
	}
}
