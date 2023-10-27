package domain

type QueueName string

const (
	HTTPSCheck QueueName = "https-check"
)

type MessageQueue interface {
	Publish(name QueueName, data []byte) error
	Consume(name QueueName) (<-chan []byte, error)
}

type noopMessageQueue struct{}

func NewNoopMessageQueue() MessageQueue {
	return &noopMessageQueue{}
}

func (n *noopMessageQueue) Publish(queueName QueueName, data []byte) error {
	// Maybe print a log or do nothing.
	return nil
}

func (n *noopMessageQueue) Consume(queueName QueueName) (<-chan []byte, error) {
	// Maybe print a log or do nothing.
	return nil, nil
}
