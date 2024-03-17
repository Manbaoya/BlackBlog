package tool

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"regexp"
)

func CheckPhone(phone string) controller.Respond {
	//检查手机号格式是否规范
	var da dao.User
	r := `^1[0-9]{10}$`
	if ok, _ := regexp.MatchString(r, phone); !ok {
		return controller.BadPhone
	}
	//检查手机号是否重复
	db := dao.DB.Where("phone=?", phone)

	result, _ := dao.Query(db, da)
	if phone == result.Phone {
		return controller.RePhone
	}
	return controller.Respond{}
}
