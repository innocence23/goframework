package middleware

import (
	"goframework/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusOK, lib.ErrorResponse(lib.NewValidTokenError("请求未携带token，无权限访问")))
			c.Abort()
			return
		}
		claims, err := lib.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, lib.ErrorResponse(lib.NewValidTokenError(err.Error())))
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
