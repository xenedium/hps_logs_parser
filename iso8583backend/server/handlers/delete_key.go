package handlers

import "github.com/gin-gonic/gin"

func DeleteKey(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Param("key")

		if key == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "Bad request"})
			return
		}

		_, err := clients.RedisClient.Do(c, "DEL", key).Result()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Key deleted", "status": "ok"})
	}
}
