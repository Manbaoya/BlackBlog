package article

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
	"time"
)

type AlterData struct {
	OldTitle   string `form:"old_title" json:"old_title" binding:"required"`
	NewTitle   string `form:"new_title" json:"new_title"`
	NewContent string `form:"new_content" json:"new_content"`
	NewTime    time.Time
	NewSort    string `form:"new_sort" json:"new_sort"`
}

func Alter(c *gin.Context) {
	var data AlterData
	err := c.ShouldBindJSON(&data)
	//err := c.ShouldBind(&data)
	if err != nil {
		log.SugaredLogger.Errorf("绑定数据失败:%s", err)
		c.JSON(200, controller.FailedBind)
		return
	}
	data.NewTime = time.Now()
	//判断标题重名
	if data.NewTitle != "" {
		db := dao.DB.Where("title=?", data.NewTitle)
		var da dao.Article
		if _, ok := dao.Query(db, da); ok {
			c.JSON(200, controller.ReTitle)
			return
		}
	}
	var result dao.Article
	var ok bool
	db := dao.DB.Where("title=?", data.OldTitle)
	result, ok = dao.Query(db, dao.Article{})
	if !ok {
		c.JSON(200, controller.FailedAlter)
		return

	}
	result = dao.Article{
		Id:      result.Id,
		Title:   data.NewTitle,
		Content: data.NewContent,
		Time:    data.NewTime,
		Author:  result.Author,
		Sort:    data.NewSort,
	}
	if ok := dao.Alter(dao.DB, result.Id, result); !ok {
		c.JSON(200, controller.FailedAlter)
		return
	}
	c.JSON(200, controller.OK)

}
