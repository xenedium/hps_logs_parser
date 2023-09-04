package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

type SSHEndpointIncomingData struct {
	Host             string `json:"host"`
	Port             uint   `json:"port"`
	User             string `json:"user"`
	Password         string `json:"password"`
	PrivateKey       string `json:"private_key"`
	Passphrase       string `json:"passphrase"`
	AbsoluteDir      string `json:"absoluteDir"`
	ParseRequestName string `json:"parseRequestName"`
}

func SSHEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var incomingData SSHEndpointIncomingData

		if err := c.ShouldBindJSON(&incomingData); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}

		if incomingData.Host == "" || incomingData.User == "" || incomingData.AbsoluteDir == "" || incomingData.ParseRequestName == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "missing required fields"})
			return
		}

		conn, err := grpc.Dial(os.Getenv("GRPC_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		defer conn.Close()

		client := protocolBuffer.NewParserClient(conn)
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		reply, err := client.SSHParse(ctx, &protocolBuffer.SSHRequest{
			Host:        incomingData.Host,
			Port:        uint64(incomingData.Port),
			User:        incomingData.User,
			Password:    &incomingData.Password,
			PrivateKey:  &incomingData.PrivateKey,
			Passphrase:  &incomingData.Passphrase,
			AbsoluteDir: incomingData.AbsoluteDir,
		})

		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": reply.Messages})
	}
}
