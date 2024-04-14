package comments

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/controller/tool"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
)

type QueryComment struct {
	ArticleId int `form:"article_id" json:"article_id"`
	UserId    int `form:"user_id" json:"user_id"`
	Page      int `form:"page" json:"page" url:"page" default:"0"`
}
type Result struct {
	Data    []dao.Comment
	Respond controller.Respond
}

func Query(c *gin.Context) {
	var data QueryComment
	if err := c.ShouldBindJSON(&data); err != nil {
		log.SugaredLogger.Errorf("绑定数据失败:%s", err)
		c.JSON(200, controller.FailedBind)
		return
	}
	limit, offset := tool.Pagination(data.Page) //基于偏移量分页
	results, ok := dao.QuerysByLimit(dao.DB, dao.Comment{}, limit, offset)
	if !ok {
		c.JSON(200, controller.FailedQuery)
	}
	c.JSON(200, Result{
		Data:    results,
		Respond: controller.OK,
	})
}
