package user

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	i, exist := c.Get("id")
	id := i.(int)
	if !exist {
		c.JSON(200, controller.FailedGetId)
		return
	}
	ok := dao.Delete(dao.DB, id, dao.User{})
	if !ok {
		c.JSON(200, controller.FailedDelete)
		return
	}
	c.JSON(200, controller.OK)
}
