package article

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateData struct {
	Title   string ` form:"title" json:"title" binding:"required"`
	Content string ` form:"content" json:"content" binding:"required"`
	Sort    string ` form:"sort" json:"sort" `
	Time    time.Time
}

func Create(c *gin.Context) {
	var data CreateData
	err := c.ShouldBindJSON(&data)
	//err := c.ShouldBind(&data)
	if err != nil {
		log.SugaredLogger.Errorf("绑定数据失败:%s", err)
		c.JSON(200, controller.FailedBind)
		return
	}
	data.Time = time.Now()
	res, exist := c.Get("username")
	username := res.(string)
	if !exist {
		c.JSON(200, controller.FailedGetUsername)
		return
	}
	//判断标题是否重名
	db := dao.DB.Where("title=?", data.Title)
	var da dao.Article
	if _, ok := dao.Query(db, da); ok {
		c.JSON(200, controller.ReTitle)
		return
	}
	var result = dao.Article{
		Title:   data.Title,
		Content: data.Content,
		Author:  username,
		Time:    data.Time,
		Sort:    data.Sort,
	}
	if _, ok := dao.Add(dao.DB, result); !ok {
		c.JSON(200, controller.FailedCreate)
		return
	}
	c.JSON(200, controller.OK)
}
