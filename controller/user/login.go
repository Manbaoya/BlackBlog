package user

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `form:"username" json:"username" binding:"required" validator:"lte=12" `
	Password string `form:"password" json:"password" binding:"required" validator:"lte=12,gte=6"`
}
type Response struct {
	Respond controller.Respond
	Token   string `json:"token"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	//err := c.ShouldBind(&data)
	if err != nil {
		log.SugaredLogger.Errorf("绑定数据失败:%s", err)
		c.JSON(200, controller.FailedBind)
		return
	}
	db := dao.DB.Where("username=?", data.Username)
	da := dao.User{
		Username: data.Username,
	}
	result, ok := dao.Query(db, da)
	if !ok {
		c.JSON(200, Response{controller.InvalidUserName, ""})
		return
	}
	//验证密码
	if ok := controller.Decrypt(result.Password, data.Password); !ok {
		c.JSON(200, Response{controller.ErrorPassword, ""})
		return
	}
	TokenString, err := controller.GenToken(result.Id, result.Username)
	if err != nil {
		log.SugaredLogger.Errorf("生成token失败:%s", err)
		c.JSON(200, Response{controller.FailedCreateToken, ""})
		return
	}
	c.JSON(200, Response{controller.OK, TokenString})

}
