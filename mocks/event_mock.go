package mock

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type PublisherStub struct {
	Messages map[string]*message.Message
}

func NewPublisherStub() *PublisherStub {
	message := make(map[string]*message.Message)
	return &PublisherStub{
		Messages: message,
	}
}

func (p *PublisherStub) Publish(topic string, messages ...*message.Message) error {
	p.Messages[topic] = messages[0]

	return nil
}

func (p *PublisherStub) Close() error {
	p.Messages = nil

	return nil
}
