package bootstrap

import (
	"fmt"

	"go.kimpton.io/url-checker/internal/config"
	"go.kimpton.io/url-checker/internal/domain"
	"go.kimpton.io/url-checker/internal/storage/rabbitmq"
)

func GetMessageQueueProvider(c *config.Config) (domain.MessageQueue, error) {
	switch c.MessageQueueProvider {
	case "rabbitmq":
		return rabbitmq.Connect(c.RabbitMQ)
	case "noop":
		return domain.NewNoopMessageQueue(), nil
	default:
		return nil, fmt.Errorf("unknown message queue provider: %s", c.MessageQueueProvider)
	}
}
