package pkg

import "context"

type Options struct {
	Payload interface{}
}

type Handler func(ctx context.Context, options Options)

type IPubSub interface {
	Publish(event string, payload interface{}) error
	Subscribe(event string, handler Handler)
}
