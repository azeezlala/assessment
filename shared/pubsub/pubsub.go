package pubsub

import "context"

const (
	ResourceAdded = "resource-added"
	CustomerAdded = "customer-added"
)

type Options struct {
	Payload interface{}
}

type Handler func(ctx context.Context, options Options)

type IPubSub interface {
	Publish(event string, payload interface{}) error
	Subscribe(event string, handler Handler)
}
