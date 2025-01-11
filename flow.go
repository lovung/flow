package flow

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Step[T context.Context] func(ctx T) error

// Seq run steps consecutively / sequentially
func Seq[T context.Context](steps ...Step[T]) Step[T] {
	return func(ctx T) error {
		for _, s := range steps {
			err := s(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

type EmbedableContext interface {
	context.Context
	Embed(context.Context)
}

// Go run steps concurrently
func Go[T EmbedableContext](steps ...Step[T]) Step[T] {
	return func(ctx T) error {
		eg, _ctx := errgroup.WithContext(ctx)
		ctx.Embed(_ctx)
		for _, s := range steps {
			eg.Go(func() error {
				return s(ctx)
			})
		}
		return eg.Wait()
	}
}
