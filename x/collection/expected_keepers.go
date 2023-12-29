package collection

import (
	"context"
)

type (
	// ClassKeeper defines the contract needed to be fulfilled for class dependencies.
	ClassKeeper interface {
		NewID(ctx context.Context) string
		HasID(ctx context.Context, id string) bool
	}
)
