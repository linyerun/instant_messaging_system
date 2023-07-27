package utils

import "github.com/gin-gonic/gin"

func DefaultPostForm(c *gin.Context, key, defaultValue string) string {
	val := c.DefaultPostForm(key, defaultValue)
	if val == "" {
		return defaultValue
	}
	return val
}
