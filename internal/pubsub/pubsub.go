package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// PublishJSON publishes a JSON-encoded message to the specified RabbitMQ exchange and routing key.
func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	// Encode the value as JSON
	body, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("failed to encode value as JSON: %w", err)
	}

	// Publish the message to the exchange with the routing key using PublishWithContext
	err = ch.PublishWithContext(
		context.Background(),
		exchange, // exchange name
		key,      // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

