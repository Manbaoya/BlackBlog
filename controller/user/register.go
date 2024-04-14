package user

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
	"regexp"
)

type RegisterData struct {
	Username   string `form:"username" json:"username"  binding:"required" validator:"lte=12"`
	Password   string `form:"password" json:"password" binding:"required" validator:"lte=12,gte=6"`
	RePassword string `form:"re_password" json:"re_password" binding:"required" validator:"lte=12,gte=6"`
	Phone      string `form:"phone" json:"phone"  validator:"len=11"`
}

func Register(c *gin.Context) {
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	//err := c.ShouldBind(&data)
	if err != nil {
		log.SugaredLogger.Errorf("绑定数据失败:%s", err)
		c.JSON(200, controller.FailedBind)
		return
	}
	//检查用户名是否规范
	r := "[^ ]{1,12}"
	if ok, _ := regexp.MatchString(r, data.Username); !ok {
		c.JSON(200, controller.BadUserName)
		return
	}
	//检查用户名是否重复
	var da dao.User
	db := dao.DB.Where("username=?", data.Username)

	result, ok := dao.Query(db, da)

	if ok {
		c.JSON(200, controller.ReUserName)
		return
	}

	//检查密码是否符合规范
	r = "[a-zA-Z0-9@#$&.]{6,12}"
	ok, _ = regexp.MatchString(r, data.Password)
	if !ok {
		c.JSON(200, controller.BadPassword)
		return
	}
	//检查两次密码是否一致
	if data.RePassword != data.Password {
		c.JSON(200, controller.DifferentPwd)
		return
	}
	if data.Phone != "" {
		//检查手机号格式是否规范
		r2 := "^1[0-9]{10}"
		ok, _ = regexp.MatchString(r2, data.Phone)
		if !ok {
			c.JSON(200, controller.BadPhone)
			return
		}
		//检查手机号是否重复
		db = dao.DB.Where("phone=?", data.Phone)

		result, ok = dao.Query(db, da)
		if data.Phone == result.Phone {
			c.JSON(200, controller.RePhone)
			return
		}
	}

	//加密密码
	data.Password, ok = controller.Encrypt(data.Password)
	if !ok {
		c.JSON(200, controller.FailedEncrypt)
		return
	}
	var re = dao.User{
		Username: data.Username,
		Password: data.Password,
		Phone:    data.Phone,
	}
	if _, ok := dao.Add(dao.DB, re); !ok {
		c.JSON(200, controller.FailedCreate)
		return
	}
	c.JSON(200, controller.OK)

}
