package middleware

import (
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/example/go-core/internal/pkg/auth"
)

func JWT(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		claims, err := auth.ParseToken(secret, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func Authorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetString("username")
		obj := c.FullPath()
		act := c.Request.Method
		allowed, err := e.Enforce(user, obj, act)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "enforce failed"})
			return
		}
		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			return
		}
		c.Next()
	}
}
