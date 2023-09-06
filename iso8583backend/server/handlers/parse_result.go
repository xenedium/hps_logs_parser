package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
	"github.com/xenedium/hps_logs_parser/iso8583backend/server/types"
	"strconv"
	"strings"
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

		responseMessages := make([]*protocolBuffer.Message, 0)

		for _, message := range parseResult.Messages {
			// if search is empty, return 0 messages
			if search.MtiVersion == "" && search.MtiClass == "" && search.MtiFunction == "" && search.MtiOrigin == "" && search.Bitmap == "" && len(search.LogFiles) == 0 && len(search.Fields) == 0 {
				break
			}

			if search.MtiVersion != "" && strconv.Itoa(int(message.Mti.Version)) != search.MtiVersion {
				continue
			}

			if search.MtiClass != "" && strconv.Itoa(int(message.Mti.Class)) != search.MtiClass {
				continue
			}

			if search.MtiFunction != "" && strconv.Itoa(int(message.Mti.Function)) != search.MtiFunction {
				continue
			}

			if search.MtiOrigin != "" && strconv.Itoa(int(message.Mti.Origin)) != search.MtiOrigin {
				continue
			}

			if search.Bitmap != "" && message.Bitmap != search.Bitmap {
				continue
			}

			if len(search.LogFiles) != 0 && !stringInSlice(message.LogFileName, search.LogFiles) {
				continue
			}

			if len(search.Fields) != 0 {
				for key, value := range search.Fields {
					fld := message.Fields[key]

					if fld == nil || !strings.Contains(fld.Value, value) {
						continue
					}
					responseMessages = append(responseMessages, message)
				}
			}
		}

		parseResult.Messages = responseMessages
		c.JSON(200, parseResult)
	}
}

func stringInSlice(name string, files []string) bool {
	for _, file := range files {
		if strings.Contains(name, file) {
			return true
		}
	}
	return false
}
