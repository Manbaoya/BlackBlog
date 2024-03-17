package dao

import (
	"BLACKBLOG/log"
	"gorm.io/gorm"
)

func Alter[t T](db *gorm.DB, id int, data t) bool {
	result := db.Model(data).Where("id=?", id).Updates(data)
	if result.Error != nil {
		log.SugaredLogger.Errorf("修改数据失败：%s", result.Error.Error())
		return false
	}
	return true
}
