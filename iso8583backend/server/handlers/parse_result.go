package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
	"github.com/xenedium/hps_logs_parser/iso8583backend/server/types"
)

/*
export interface Search {
    mtiVersion?: string;
    mtiClass?: string;
    mtiFunction?: string;
    mtiOrigin?: string;
    bitmap?: string;
    logFiles?: string[];
    fields?: { [key: string]: string };
}
*/

type Search struct {
	MtiVersion  string            `json:"mtiVersion"`
	MtiClass    string            `json:"mtiClass"`
	MtiFunction string            `json:"mtiFunction"`
	MtiOrigin   string            `json:"mtiOrigin"`
	Bitmap      string            `json:"bitmap"`
	LogFiles    []string          `json:"logFiles"`
	Fields      map[string]string `json:"fields"`
}

func GetParseResult(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		var search Search
		if err := c.ShouldBindJSON(&search); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		}

		key := c.Param("key")

		result, err := clients.RedisClient.Do(c, "GET", key).Result()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		if result == nil {
			c.AbortWithStatusJSON(404, gin.H{"error": "not found"})
			return
		}

		parseResult := &types.ParseResult{}

		if err := json.Unmarshal([]byte(result.(string)), parseResult); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		parseResult.Messages = make([]*protocolBuffer.Message, 0)

		c.JSON(200, parseResult)
	}
}
