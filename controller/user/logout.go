package user

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Logout(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(200, controller.EmptyAuth)
		c.Abort()
		return
	}
	//按空格分割
	parts := strings.SplitN(authHeader, " ", 2)

	dao.RDB.SAdd(dao.Ctx, "blacklist", parts[1])
	es, err1 := dao.RDB.SMembers(dao.Ctx, "blacklist").Result()
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(es)
	c.JSON(200, "成功")
}
