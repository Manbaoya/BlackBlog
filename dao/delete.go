package dao

import (
	"BLACKBLOG/log"
	"gorm.io/gorm"
)

func Delete[t T](db *gorm.DB, id int, data t) bool {
	result := db.Model(data).Where("id=?", id).Delete(data)
	if result.Error != nil {
		log.SugaredLogger.Errorf("删除数据失败：%s", result.Error.Error())
		return false
	}
	return true
}
