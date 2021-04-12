package web

import (
	"github.com/Gateway/pkg/webserver"
)

type Web struct {
	Client *webserver.WebServer
}

func Init(port string) (*Web, error) {
	client, err := webserver.Init(port)
	if err != nil {
		return nil, err
	}
	return &Web{
		Client: client,
	}, nil
}
