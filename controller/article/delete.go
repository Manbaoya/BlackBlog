package article

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
)

type DeleteData struct {
	Id int `form:"id" json:"id"`
}

func Delete(c *gin.Context) {
	var data DeleteData
	err := c.ShouldBindJSON(&data)
	//err := c.ShouldBind(&data)
	if err != nil {
		log.SugaredLogger.Errorf("绑定数据失败:%s", err)
		c.JSON(200, controller.FailedBind)
		return
	}

	ok := dao.Delete(dao.DB, data.Id, dao.Article{})
	if !ok {
		c.JSON(200, controller.FailedDelete)
		return
	}
	c.JSON(200, controller.OK)
}
