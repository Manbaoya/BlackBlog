package tool

import (
	"BLACKBLOG/controller"
	"regexp"
)

func CheckPassword(password, rePassword string) controller.Respond {
	//检查密码是否符合规范
	r := `^[a-zA-Z0-9*&$#.]{6,12}$`
	ok, _ := regexp.MatchString(r, password)
	if !ok {
		return controller.BadPassword
	}
	//检查两次密码是否一致
	if rePassword != password {
		return controller.DifferentPwd
	}
	return controller.Respond{}
}
