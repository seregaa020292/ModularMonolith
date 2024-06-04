package decorator

import "context"

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}
