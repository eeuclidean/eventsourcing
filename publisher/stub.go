package publisher

import (
	"errors"

	"github.com/eeuclidean/eventsourcing"
)

type PublisherSuccessStub struct{}

func (mock *PublisherSuccessStub) Publish(event eventsourcing.Event) error {
	return nil
}

type PublisherErrorStub struct{}

func (mock *PublisherErrorStub) Publish(event eventsourcing.Event) error {
	return errors.New("error")
}
