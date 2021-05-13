package common

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	TicketCreated = "ticket-created"
	TIcketUpdated = "ticket-updated"

	OrderCreated   = "order-created"
	OrderCancelled = "order-cancelled"

	ExpirationComplete = "expiration-complete"

	PaymentCreated = "payment-created"
)

type EventHandler func(msg *message.Message) error

func NewPublisher(brokers []string, logger watermill.LoggerAdapter) (message.Publisher, error) {
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

type SubscriberConfig struct {
	Brokers       []string
	ConsumerGroup string
	FromBeginning bool
	watermill.LoggerAdapter
}

func NewSubscriber(config *SubscriberConfig) (message.Subscriber, error) {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	if config.FromBeginning {
		saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	kafkaSubscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               config.Brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         config.ConsumerGroup,
		},
		config.LoggerAdapter,
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
			log.Println(err)
		}
	}
}
