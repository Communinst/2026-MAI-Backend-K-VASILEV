package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RequireMethods — аналог @require_http_methods в Django
// Ограничивает доступ только к указанным HTTP-методам
func RequireMethods(allowed ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, method := range allowed {
			if c.Request.Method == method {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Method not allowed",
		})
	}
}
