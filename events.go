package common

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	TicketCreated = "ticket:created"
	TIcketUpdated = "ticket:updated"

	OrderCreated   = "order:created"
	OrderCancelled = "order:cancelled"

	ExpirationComplete = "expiration:complete"

	PaymentCreated = "payment:created"
)

type EventHandler func(msg *message.Message)

func CreatePublisher(brokers []string) (message.Publisher, error) {
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   brokers,
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}

	return publisher, nil
}

func CreateSubscriber(brokers []string, consumerGroup string) (message.Subscriber, error) {
	kafkaSubscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: kafka.DefaultSaramaSubscriberConfig(),
			ConsumerGroup:         consumerGroup,
		},
		watermill.NewStdLogger(
			true,
			false,
		),
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
	go func() {
		for msg := range messages {
			eventHandler(msg)
		}
	}()

	return nil
}
