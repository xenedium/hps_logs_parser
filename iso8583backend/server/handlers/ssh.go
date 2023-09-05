package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
	"github.com/xenedium/hps_logs_parser/iso8583backend/server/types"
	"time"
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

func SSHEndpoint(clients *Clients) gin.HandlerFunc {
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

		reply, err := clients.GrpcClient.SSHParse(clients.GrpcContext, &protocolBuffer.SSHRequest{
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

		parseResult := types.ParseResult{
			Id:       incomingData.ParseRequestName,
			Name:     incomingData.ParseRequestName,
			Date:     time.Now(),
			Status:   "done",
			Type:     "ssh",
			Messages: reply.Messages,
		}

		data, err := json.Marshal(parseResult)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		if err := clients.RedisClient.Set(clients.RedisContext, incomingData.ParseRequestName, data, 0).Err(); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "success"})
	}
}
