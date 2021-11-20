package serializer

import (
	"context"
)

type Serializer interface {
	Encode(ctx context.Context, data interface{}) ([]byte, error)
	Decode(ctx context.Context, data []byte, res interface{}) error
}
