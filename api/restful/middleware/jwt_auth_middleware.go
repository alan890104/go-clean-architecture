package middleware

import (
	"net/http"
	"strings"

	"github.com/alan890104/go-clean-arch-demo/domain"
	tokensvc "github.com/alan890104/go-clean-arch-demo/internal/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		bearerToken := strings.Split(authHeader, "")
		if len(bearerToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.Response{
				Msg: "Invalid token format (length different from 2)",
			})
			return
		}
		authToken := bearerToken[1]
		identity, err := tokensvc.ExtractIdentityFromToken(authToken, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.Response{
				Msg: err.Error(),
			})
			return
		}
		c.Set("identity", identity)
		c.Next()
	}
}
