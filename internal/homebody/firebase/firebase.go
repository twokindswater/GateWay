package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	_ "firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func Init(ctx context.Context, config Config) (*firebase.App, error) {
	opt := option.WithCredentialsFile(config.ConfigFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
