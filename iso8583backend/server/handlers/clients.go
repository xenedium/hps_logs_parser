package handlers

import (
	"context"
	"github.com/redis/go-redis/v9"
	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
)

type Clients struct {
	RedisClient  *redis.Client
	GrpcClient   protocolBuffer.ParserClient
	RedisContext context.Context
	GrpcContext  context.Context
}
