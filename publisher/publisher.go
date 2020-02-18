package publisher

import "github.com/eeuclidean/eventsourcing"

type DomainEventPublisher interface {
	Publish(event eventsourcing.Event) error
}
