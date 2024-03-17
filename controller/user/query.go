package user

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"github.com/gin-gonic/gin"
)

type GetQuery struct {
	Respond  controller.Respond `json:"Respond"`
	Username string             `json:"username"`
	Phone    string             `json:"phone"`
}

func Query(c *gin.Context) {
	i, exist := c.Get("id")
	id := i.(int)
	if !exist {
		c.JSON(200, controller.FailedGetId)
		return
	}
	db := dao.DB.Where("id=?", id)
	result, ok := dao.Query(db, dao.User{})
	if !ok {
		c.JSON(200, GetQuery{controller.FailedQuery, "", ""})
		return
	}
	c.JSON(200, GetQuery{controller.OK, result.Username, result.Phone})
}
