package handlers

import "github.com/gin-gonic/gin"

func GetKeys(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		keys, err := clients.RedisClient.Do(c, "KEYS", "*").Result()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"keys": keys})
	}
}
