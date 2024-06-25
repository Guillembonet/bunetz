package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AssetsCache(c *gin.Context) {
	if strings.HasPrefix(c.Request.URL.Path, "/assets/") || strings.HasPrefix(c.Request.URL.Path, "/blog/assets/") {
		c.Header("Cache-Control", "private, max-age=86400")
	}
	c.Next()

}
