package flow

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Func[T context.Context] func(ctx T) error

// Steps run steps consecutively
func Steps[T context.Context](ctx T, steps ...Func[T]) error {
	for _, s := range steps {
		err := s(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// Go run steps concurrently
func Go[T EmbedableContext](ctx T, steps ...Func[T]) error {
	eg, _ctx := errgroup.WithContext(ctx)
	ctx.Embed(_ctx)
	for _, s := range steps {
		eg.Go(func() error {
			return s(ctx)
		})
	}
	return eg.Wait()
}
