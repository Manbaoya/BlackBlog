package comments

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	i := c.Query("comment_id")
	id, err := strconv.Atoi(i)
	if err != nil {
		log.SugaredLogger.Errorf("绑定数据失败:%s", err)
		c.JSON(200, controller.FailedBind)
		return
	}
	ok := dao.Delete(dao.DB, id, dao.Comment{})
	if !ok {
		c.JSON(200, controller.FailedDelete)
		return
	}
	c.JSON(200, controller.OK)
}
