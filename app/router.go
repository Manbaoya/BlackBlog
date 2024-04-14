package app

import (
	"BLACKBLOG/app/api/middleware"
	"BLACKBLOG/controller/article"
	"BLACKBLOG/controller/comments"
	"BLACKBLOG/controller/user"
	_ "BLACKBLOG/docs"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(log.GinLogger(log.Logger), log.Recovery(log.Logger, true))
	r.StaticFS("/image", http.Dir("upload/upload"))
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/ul") //非限制
	{
		v1.POST("/login", middleware.Cors(), user.Login) //登录
		v1.POST("/register", user.Register)              //注册

	}
	v2 := r.Group("/l") //限制
	v2.Use(middleware.JwtAuthMiddleware())
	{
		v2.POST("/alter_u", user.Alter)           //修改用户信息
		v2.POST("/uploadImage", user.UploadImage) //上传头像
		v2.DELETE("delete_u", user.Delete)        //注销
		v2.GET("logout", user.Logout)             //退出登录
		v2.GET("/query_u", user.Query)            //查询信息
		v2.POST("/create_a", article.Create)      //创建文章
		v2.POST("/alter_a", article.Alter)        //修改文章
		v2.DELETE("/delete_a", article.Delete)    //删除文章
		v2.GET("/query_a", article.Query)         //查询文章

		v2.GET("/create_c", comments.Create)    //创建评论
		v2.DELETE("/delete_c", comments.Delete) //删除评论
		v2.GET("/query_c", comments.Query)      //查询评论

	}
	if err := r.Run(":8080"); err != nil {
		log.SugaredLogger.Errorf("路由启动失败:%s", err)
		return
	}
}
