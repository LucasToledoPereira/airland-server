package middlewares

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

func CheckUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		id := token.RegisteredClaims.Subject
		c.Set("UserID", id)
	}
}
