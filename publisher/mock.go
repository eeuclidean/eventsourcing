package publisher

import (
	"github.com/eeuclidean/eventsourcing"
	"github.com/stretchr/testify/mock"
)

type PublisherMock struct {
	mock.Mock
}

func (mock *PublisherMock) Publish(event eventsourcing.Event) error {
	args := mock.Called(event)
	return args.Error(0)
}
