package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

func UploadFilesEndpoint() gin.HandlerFunc {
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

		if len(parseRequestName) == 0 {
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
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

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

		reply, err := client.FilesParse(ctx, filesRequest)

		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": reply.Messages})
	}
}
