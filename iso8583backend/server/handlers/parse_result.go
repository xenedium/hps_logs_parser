package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetParseResult(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		c.Header("Content-Type", "application/json")
		c.String(200, result.(string))
	}
}
