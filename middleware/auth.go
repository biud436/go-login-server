// middleware/auth.go
package middleware

import (
	"go-login-server/auth"
	"go-login-server/common"

	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            common.ErrorResponse(c, common.ErrAuthorizationHeader)
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            common.ErrorResponse(c, common.ErrInvalidAuthHeader)
            c.Abort()
            return
        }

        tokenStr := parts[1]
        claims, err := auth.ValidateJWT(tokenStr)
        if err != nil {
            common.ErrorResponse(c, common.ErrInvalidToken)
            c.Abort()
            return
        }

        // 사용자 이름을 컨텍스트에 저장
        c.Set("username", claims.Username)
        c.Next()
    }
}
