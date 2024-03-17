package dao

import (
	"BLACKBLOG/log"
	"gorm.io/gorm"
)

func Query[t T](db *gorm.DB, data t) (t, bool) { //检索单个对象

	result := db.Model(&data).First(&data)
	if result.Error != nil {
		log.SugaredLogger.Errorf("查询数据失败：%s", result.Error.Error())
		return data, false
	}
	return data, true
}

func Querys[t T, s S](db *gorm.DB, data t, limit int, offset int) (s, bool) { //检索多个对象
	var datas []Article
	results := db.Model(&data).Limit(limit).Offset(offset).Find(&datas)
	if results.Error != nil {
		log.SugaredLogger.Errorf("查询数据失败：%s", results.Error.Error())
		return datas, false
	}
	return datas, true
}
