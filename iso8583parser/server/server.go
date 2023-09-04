package server

import (
	"flag"
	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type RestServer struct {
	address string
	router  *gin.Engine
	apiKey  string
}

func (s *RestServer) Run() {
	err := s.router.Run(s.address)
	if err != nil {
		return
	}
}

func NewRestServer(Address string, ApiKey string) *RestServer {
	newServer := &RestServer{address: Address, router: gin.Default(), apiKey: ApiKey}

	// MIDDLEWARES
	newServer.router.Use(gin.Recovery())
	newServer.router.Use(ApiKeyAuthMiddleware(ApiKey))

	// ROUTES
	v1 := newServer.router.Group("/api/v1")
	{
		v1.POST("/upload", UploadFilesEndpoint())
		v1.POST("/ssh", SSHEndpoint())
	}

	// error 404
	newServer.router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "not found"})
	})

	return newServer
}

type gRPCServer struct {
	protocolBuffer.UnimplementedParserServer
}

func NewGRPCServer(Address string) {
	flag.Parse()
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*1024),
		grpc.MaxSendMsgSize(1024*1024*1024),
	)
	protocolBuffer.RegisterParserServer(grpcServer, &gRPCServer{})

	log.Printf("Server listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
