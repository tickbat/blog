package jwt

import (
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		token := c.GetHeader("token")
		if token == "" {
			code = e.MISS_AUTH_TOKEN
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR
				}
			}
		}

		if code != e.SUCCESS {
			util.Res(c, http.StatusUnauthorized, code, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
