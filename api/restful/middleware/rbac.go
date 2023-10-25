package middleware

import (
	"net/http"
	"strings"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/alan890104/go-clean-arch-demo/rbac"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinRBACMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("identity")
		identity := &domain.Identity{}
		if exists {
			identity = id.(*domain.Identity)
		}

		role := identity.Role
		userId := identity.UserID
		obj := c.Request.URL.Path
		act := c.Request.Method

		// Remove /api/v1 prefix
		obj = strings.TrimPrefix(obj, "/api/v1")

		// Special handling all routes having userID param
		userIdParam := c.Param("userID")
		if userIdParam != "" && role == rbac.Visitor && userId != userIdParam {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		// Check permission with e.Enforce(role, obj, act)
		allowed, err := e.Enforce(role, obj, act)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
