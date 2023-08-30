package chanresult

import "context"

func sendUntilDone[T any](
	ctx context.Context,
	chn chan<- T,
	t T,
) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case chn <- t:
		return nil
	}
}
