package webserver

import "github.com/gin-gonic/gin"

type WebServer struct {
	Router *gin.Engine
	port   string
}

func Init(port string) (*WebServer, error) {
	// initialize gin server
	// TODO: setting server more detail
	return &WebServer{
		Router: gin.Default(),
		port:   port,
	}, nil
}

func (s *WebServer) Run() error {
	return s.Router.Run(s.port)
}
