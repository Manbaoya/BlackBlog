package tool

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"regexp"
)

func CheckUsername(username string) controller.Respond {
	var da dao.User
	//检查用户名是否规范
	r := `^[^ ]{1,12}$`
	if ok, _ := regexp.MatchString(r, username); !ok {
		return controller.BadUserName
	}
	//检查用户名是否重复

	db := dao.DB.Where("username=?", username)
	_, ok := dao.Query(db, da)

	if ok {
		return controller.ReUserName
	}
	return controller.Respond{}

}
