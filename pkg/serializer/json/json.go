package json

import (
	"context"
	"encoding/json"
)

type JsonSerializer struct {
}

func JsonSerializerInit() *JsonSerializer {
	return &JsonSerializer{}
}

func (j *JsonSerializer) Encode(ctx context.Context, data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (j *JsonSerializer) Decode(ctx context.Context, data []byte, res interface{}) error {
	return json.Unmarshal(data, res)
}
