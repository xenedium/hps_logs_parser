package server

import (
	"context"
	"log"

	protocolBuffer "github.com/xenedium/hps_logs_parser/server/gRPC"
)

// SSHParse TODO: implement
func (s *gRPCServer) SSHParse(ctx context.Context, in *protocolBuffer.SSHRequest) (*protocolBuffer.Response, error) {
	log.Printf("Host: %v", in.Host)
	return &protocolBuffer.Response{
		RequestId: "2",
	}, nil
}
