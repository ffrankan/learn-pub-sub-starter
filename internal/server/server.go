package server

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

// PublishJSON publishes a message to the specified exchange with the given routing key
func PublishJSON(ch *amqp.Channel, exchange, routingKey string, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
} 