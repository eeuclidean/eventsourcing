package eventstore

import "github.com/eeuclidean/eventsourcing"

type EventStore interface {
	Save(event eventstorming.Event) error
}
