package jwt

import (
	"apisix-admin/consts/error_code"
	"apisix-admin/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		notCheckPath := []string{
			"/apisix_admin/v1/organization/user/login",
		}
		for _, path := range notCheckPath {
			if path == c.Request.URL.Path {
				return
			}
		}
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": error_code.CODE_AUTH_TOKEN_EXPIRE,
				"message":  error_code.ErrCodeMessage(error_code.CODE_AUTH_TOKEN_EXPIRE),
				"data": nil,
			})
			c.Abort()
			return
		}
		ok, err := jwt.CheckToken(token)
		if err == nil && ok {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": error_code.CODE_AUTH_TOKEN_EXPIRE,
				"message":  error_code.ErrCodeMessage(error_code.CODE_AUTH_TOKEN_EXPIRE),
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
