package serializer

import (
	"context"
	"errors"
)
import "github.com/HomeLongServer/pkg/serializer/json"

var (
	errUndefinedSerializerType = errors.New("undefined serializer type or wrong serializer type")
)

const (
	Json = "json"
)

type Serializer interface {
	Encode(ctx context.Context, data interface{}) ([]byte, error)
	Decode(ctx context.Context, data []byte, res interface{}) error
}

func Init(sType string) (Serializer, error) {
	switch sType {
	case Json:
		return json.JsonSerializerInit(), nil
	default:
		return nil, errUndefinedSerializerType
	}
}
