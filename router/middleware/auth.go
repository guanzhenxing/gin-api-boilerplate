package middleware

import (
	"gin-api-boilerplate/handler"
	"gin-api-boilerplate/pkg/ecode"
	"gin-api-boilerplate/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, ecode.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
