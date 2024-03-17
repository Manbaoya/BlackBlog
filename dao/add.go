package dao

import (
	"BLACKBLOG/log"
	"gorm.io/gorm"
)

func Add[t T](db *gorm.DB, data t) bool {
	result := db.Model(&data).Create(&data)
	if result.Error != nil {
		log.SugaredLogger.Errorf("新增数据失败:%s", result.Error.Error())
		return false
	}
	return true
}
