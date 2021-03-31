package json

import "context"

type JsonSerializer struct {
}

func JsonSerializerInit() *JsonSerializer {
	return &JsonSerializer{}
}

func (j *JsonSerializer) Encode(ctx context.Context, data interface{}) ([]byte, error) {
	return nil, nil
}

func (j *JsonSerializer) Decode(ctx context.Context, data []byte, res interface{}) error {
	return nil
}
