package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	maxAge = 12
)

//	func Cors() gin.HandlerFunc {
//		return cors.New(cors.Config{
//			AllowOrigins:     []string{"*"},
//			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
//			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
//			ExposeHeaders:    []string{"Content-Length"},
//			AllowCredentials: true,
//			AllowOriginFunc: func(origin string) bool {
//				return origin == "https://github.com"
//			},
//			MaxAge: maxAge * time.Hour,
//		})
//
// }
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,token,authorizatiowsn")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
