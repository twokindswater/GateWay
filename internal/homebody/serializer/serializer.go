package serializer

import (
	"context"
	"errors"
	"github.com/Gateway/internal/homebody/serializer/json"
	"github.com/Gateway/pkg/serializer"
)

var (
	errUndefinedSerializerType = errors.New("undefined serializer type or wrong serializer type")
)

const (
	Json = "json"
)

func Init(ctx context.Context, config Config) (serializer.Serializer, error) {
	switch config.Type {
	case Json:
		return json.JsonSerializerInit(), nil
	default:
		return nil, errUndefinedSerializerType
	}
}
