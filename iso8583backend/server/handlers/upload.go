package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
)

func UploadFilesEndpoint(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}

		parseRequestName := form.Value["parseRequestName"]
		files := form.File["files"]

		if len(files) == 0 {
			c.AbortWithStatusJSON(400, gin.H{"error": "no files uploaded"})
			return
		}

		if len(parseRequestName) != 1 || parseRequestName[0] == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "parseRequestName is required and must be unique"})
			return
		}

		var filesRequest = &protocolBuffer.FilesRequest{Files: []*protocolBuffer.File{}}

		for _, file := range files {
			openedFile, err := file.Open()
			if err != nil {
				c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
				return
			}
			fileData := make([]byte, file.Size)
			_, err = openedFile.Read(fileData)
			if err != nil {
				c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
				return
			}
			filesRequest.Files = append(filesRequest.Files, &protocolBuffer.File{
				Name:    file.Filename,
				Content: string(fileData),
			})
		}

		reply, err := clients.GrpcClient.FilesParse(clients.GrpcContext, filesRequest)

		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		data, err := json.Marshal(reply.Messages)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		if err := clients.RedisClient.Set(clients.RedisContext, parseRequestName[0], data, 0).Err(); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "success"})
	}
}
