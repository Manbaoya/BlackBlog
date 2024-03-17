package middleware

import (
	"BLACKBLOG/controller"
	"github.com/gin-gonic/gin"
	"strings"
)

//检验

func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("authorizatiowsn")
		if authHeader == "" {
			c.JSON(200, controller.EmptyAuth)
			c.Abort()
			return
		}
		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(200, controller.ErrorAuth)
			c.Abort()
			return
		}
		mc, err := controller.ParseToken(parts[1])
		if err != nil {
			c.JSON(200, controller.InvalidToken)
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Set("id", mc.Id)
		c.Next()
	}

}
