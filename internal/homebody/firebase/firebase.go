package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	_ "firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"

	"google.golang.org/api/option"
)

var (
	fcm_repository = "../../fcm_keystore/"
)

func Init(ctx context.Context, config Config) (*messaging.Client, error) {
	opt := option.WithCredentialsFile(fcm_repository + "homebody-311609-firebase-adminsdk-ae9o9-35132d50bc.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("error messaping app: %v", err)
	}

	return client, nil
}
