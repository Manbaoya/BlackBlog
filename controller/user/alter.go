package user

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/controller/user/tool"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
	"reflect"
)

type AlterData struct {
	NewUsername string `form:"new_username" json:"new_username"`
	NewPassword string `form:"new_password" json:"new_password"`
	RePassword  string `form:"re_password" json:"re_password"`
	NewPhone    string `form:"new_phone" json:"new_phone"`
}

func Alter(c *gin.Context) {
	var data AlterData
	err := c.ShouldBindJSON(&data)
	//err := c.ShouldBind(&data)
	if err != nil {
		log.SugaredLogger.Errorf("绑定数据失败:%s", err)
		c.JSON(200, controller.FailedBind)
		return
	}
	i, exist := c.Get("id")
	id := i.(int)
	if !exist {
		c.JSON(200, controller.FailedGetId)
		return
	}

	//检查用户名规范重复
	if data.NewUsername != "" {
		res := tool.CheckUsername(data.NewUsername)
		if !reflect.DeepEqual(res, controller.Respond{}) {
			c.JSON(200, Response{res, ""})
			return
		}
	}

	//检查密码规范和一致
	if data.NewPassword != "" {
		res := tool.CheckPassword(data.NewPassword, data.RePassword)
		if !reflect.DeepEqual(res, controller.Respond{}) {
			c.JSON(200, Response{res, ""})
			return
		}
		//密码加密
		var ok bool
		data.NewPassword, ok = controller.Encrypt(data.NewPassword)
		if !ok {
			c.JSON(200, Response{controller.FailedEncrypt, ""})
			return
		}
	}
	//检查手机号的规范和重复
	if data.NewPhone != "" {
		res := tool.CheckPhone(data.NewPhone)
		if !reflect.DeepEqual(res, controller.Respond{}) {
			c.JSON(200, Response{res, ""})
			return
		}
	}
	//更新数据库数据
	var result = dao.User{
		Username: data.NewUsername,
		Password: data.NewPassword,
		Phone:    data.NewPhone,
	}
	dao.Alter(dao.DB, id, result)
	//更新token
	token, err := controller.GenToken(id, data.NewUsername)
	if err != nil {
		log.SugaredLogger.Errorf("生成token失败:%s", err)
		c.JSON(200, gin.H{
			"code":    controller.FailedCreateToken.Code,
			"message": controller.FailedCreateToken.Message,
			"token":   "",
		})
		return
	}
	c.JSON(200, Response{controller.OK, token})

}
