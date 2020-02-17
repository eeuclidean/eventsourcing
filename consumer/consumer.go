package consumer

import (
	"context"
	"github.com/eeuclidean/eventsourcing"
)

type EventConsumer interface {
	Run(ctx context.Context) error
}

type EventConsumerHandler interface {
	Apply(event eventstorming.Event) error
}
