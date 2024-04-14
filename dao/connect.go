package dao

import (
	. "BLACKBLOG/config"
	"BLACKBLOG/log"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client
var Ctx = context.Background()

func Connect() {
	dsn := Conf.DataBase.User + ":" + Conf.DataBase.Password + "@tcp(" + Conf.Server.Address + ":" + Conf.DataBase.Port + ")/" + Conf.DataBase.Name + "?charset=" + Conf.DataBase.Charset + "&parseTime=" + Conf.DataBase.ParseTime + "&loc=" + Conf.DataBase.Loc
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.SugaredLogger.Errorf("数据库连接失败: %s", err.Error())
		return
	}
	fmt.Println("数据库连接成功")
	RedisConnect()
	fmt.Println("redis连接成功")
}

func RedisConnect() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     Conf.Redis.Addr,
		Password: Conf.Redis.Password,
		DB:       Conf.Redis.DB,
		PoolSize: Conf.Redis.PoolSize,
	})

}
