package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
	"github.com/xenedium/hps_logs_parser/iso8583backend/server/handlers"
	"github.com/xenedium/hps_logs_parser/iso8583backend/server/middlewares"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

type Server struct {
	address string
	router  *gin.Engine
}

func (s *Server) Run() {
	err := s.router.Run(s.address)
	if err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}

func NewServer(Address string) *Server {
	clients := &handlers.Clients{}
	// REDIS
	clients.RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	clients.RedisContext = context.Background()
	if _, err := clients.RedisClient.Ping(clients.RedisContext).Result(); err != nil {
		log.Fatalf("Error connecting to Redis server: %v", err)
	}
	log.Printf("Connected to Redis server: %v", os.Getenv("REDIS_ADDRESS"))

	// GRPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	conn, err := grpc.DialContext(ctx, os.Getenv("GRPC_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error connecting to GRPC server: %v", err)
	}
	defer cancel()
	log.Printf("Connected to GRPC server: %v", os.Getenv("GRPC_ADDRESS"))

	clients.GrpcClient = protocolBuffer.NewParserClient(conn)
	clients.GrpcContext = context.WithoutCancel(context.Background())

	newServer := &Server{address: Address, router: gin.Default()}

	// MIDDLEWARES
	newServer.router.Use(middlewares.CORSMiddleware())

	// ROUTES
	v1 := newServer.router.Group("/api/v1")
	{
		v1.POST("/upload", handlers.UploadFilesEndpoint(clients))
		v1.POST("/ssh", handlers.SSHEndpoint(clients))
		v1.GET("/keys", handlers.GetKeys(clients))
		v1.GET("/parse/:key", handlers.GetParseResult(clients))
	}

	// error 404
	newServer.router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "not found"})
	})

	return newServer
}
