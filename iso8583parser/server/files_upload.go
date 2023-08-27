package server

import (
	"github.com/gin-gonic/gin"
	"github.com/xenedium/iso8583parser/parser"
	"os"
	"path"
)

func UploadFilesEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}

		files := form.File["files"]

		if len(files) == 0 {
			c.AbortWithStatusJSON(400, gin.H{"error": "no files uploaded"})
			return
		}

		tempDir, err := os.MkdirTemp(os.TempDir(), "*")
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		for _, file := range files {
			err := c.SaveUploadedFile(file, path.Join(tempDir, file.Filename))
			if err != nil {
				c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
				return
			}
		}

		logParser := parser.NewParser(tempDir)
		logParser.Parse(true)
		c.JSON(200, gin.H{"messages": logParser.Messages})
	}
}
