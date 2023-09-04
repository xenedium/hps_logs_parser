package server

import (
	"github.com/gin-gonic/gin"
	"github.com/xenedium/hps_logs_parser/iso8583backend/server/handlers"
)

type Server struct {
	address string
	router  *gin.Engine
}

func (s *Server) Run() {
	err := s.router.Run(s.address)
	if err != nil {
		return
	}
}

func NewServer(Address string) *Server {
	newServer := &Server{address: Address, router: gin.Default()}

	// MIDDLEWARES
	newServer.router.Use(gin.Recovery())

	// ROUTES
	v1 := newServer.router.Group("/api/v1")
	{
		v1.POST("/upload", handlers.UploadFilesEndpoint())
		v1.POST("/ssh", handlers.SSHEndpoint())
	}

	// error 404
	newServer.router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "not found"})
	})

	return newServer
}
