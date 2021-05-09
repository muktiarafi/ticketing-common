package common

import (
	"context"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	TicketCreated = "ticket.created"
	TIcketUpdated = "ticket.updated"

	OrderCreated   = "order.created"
	OrderCancelled = "order.cancelled"

	ExpirationComplete = "expiration.complete"

	PaymentCreated = "payment.created"
)

type (
	TicketCreatedEvent struct {
		ID      int     `msgpack:"id"`
		Version int     `msgpack:"version"`
		Title   string  `msgpack:"title"`
		Price   float64 `msgpack:"price"`
		UserID  int     `msgpack:"userId"`
	}

	TicketUpdatedEvent struct {
		ID      int     `msgpack:"id"`
		Version int     `msgpack:"version"`
		Title   string  `msgpack:"title"`
		Price   float64 `msgpack:"price"`
		UserID  int     `msgpack:"userId"`
		OrderID int     `msgpack:"orderId"`
	}

	OrderCreatedEvent struct {
		ID          int     `msgpack:"id"`
		Status      string  `msgpack:"status"`
		Version     int     `msgpack:"version"`
		UserID      int     `msgpack:"userId"`
		ExpiresAt   string  `msgpack:"expiresAt"`
		TicketID    int     `msgpack:"ticketId"`
		TicketPrice float64 `msgpack:"ticketPrice"`
	}

	OrderCancelledEvent struct {
		ID       int `msgpack:"id"`
		Version  int `msgpack:"version"`
		TicketID int `msgpack:"ticketId"`
	}
)

type EventHandler func(msg *message.Message) error

func CreatePublisher(brokers []string, logger watermill.LoggerAdapter) (message.Publisher, error) {
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   brokers,
			Marshaler: kafka.DefaultMarshaler{},
		},
		logger,
	)
	if err != nil {
		return nil, err
	}

	return publisher, nil
}

func CreateSubscriber(brokers []string, consumerGroup string, logger watermill.LoggerAdapter) (message.Subscriber, error) {
	kafkaSubscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: kafka.DefaultSaramaSubscriberConfig(),
			ConsumerGroup:         consumerGroup,
		},
		logger,
	)
	if err != nil {
		return nil, err
	}

	return kafkaSubscriber, nil
}

type Consumer struct {
	message.Subscriber
}

func NewConsumer(subscriber message.Subscriber) *Consumer {
	return &Consumer{
		Subscriber: subscriber,
	}
}

func (c *Consumer) On(topic string, eventHandler EventHandler) error {
	messages, err := c.Subscribe(context.Background(), topic)
	if err != nil {
		return err
	}
	go process(messages, eventHandler)

	return nil
}

func process(messages <-chan *message.Message, eventHandler EventHandler) {
	for msg := range messages {
		if err := eventHandler(msg); err != nil {
			log.Fatal(err)
		}
	}
}
