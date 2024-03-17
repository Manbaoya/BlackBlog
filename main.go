package main

import (
	"BLACKBLOG/app"
	"BLACKBLOG/config"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
)

//	@title		BLACKBLOG接口文档
//	@version	1.0
//	@description
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	满宝呀
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		8080
//	@BasePath	localhost:8080
func main() {

	config.GetConfig() //获取全局配置
	log.InitLog()      //日志初始化
	dao.Connect()      //
	log.SugaredLogger.Errorf("撞大运咯")
	app.InitRouter() //路由初始化

}
