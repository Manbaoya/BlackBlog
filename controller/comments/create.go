package comments

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type CreateComment struct {
	ArticleId  int    `form:"article_id" json:"article_id"`
	UserId     int    `form:"user_id" json:"user_id"`
	Content    string `form:"content" json:"content"`
	CreateTime time.Time
}

func Create(c *gin.Context) {
	var data CreateComment
	if err := c.ShouldBindJSON(&data); err != nil {
		log.SugaredLogger.Errorf("绑定数据失败:%s", err)
		c.JSON(200, controller.FailedBind)
		return
	}
	userId, ok := c.Get("id")
	if !ok {
		c.JSON(200, controller.FailedGetId)
		return
	}
	data.UserId = userId.(int)
	data.CreateTime = time.Now()
	da := dao.Comment{
		ArticleId:  data.ArticleId,
		UserId:     data.UserId,
		Content:    data.Content,
		CreateTime: data.CreateTime}

	result, ok := dao.Add(dao.DB, da)

	//dao.RDB.Set(dao.Ctx, "CommentSet",strconv.Itoa(result.Id),0) //放在评论id集合里
	//dao.RDB.Set(dao.Ctx,"Article",strconv.Itoa(result.ArticleId),0)  //存储文章在文章集合里
	dao.RDB.Set(dao.Ctx, "Comment"+strconv.Itoa(result.Id), result, 0) //存评论的k-v
	length, err := dao.RDB.LLen(dao.Ctx, "Article"+strconv.Itoa(result.ArticleId)).Result()
	fmt.Println("长度", length)
	if err != nil {
		return
	}
	if length >= 20 {
		x := dao.RDB.LPop(dao.Ctx, "Article"+strconv.Itoa(result.ArticleId)) //删掉20条评论中最早的一条评论
		fmt.Println("被删评论:", x)
	}
	dao.RDB.RPush(dao.Ctx, "Article"+strconv.Itoa(result.ArticleId), strconv.Itoa(result.Id)) //存放到最新评论列表里
	if !ok {
		c.JSON(200, controller.FailedCreate)
		return
	}
	Vals, err := dao.RDB.LRange(dao.Ctx, "Article"+strconv.Itoa(result.ArticleId), 0, -1).Result()
	fmt.Println(Vals)
	c.JSON(200, controller.OK)

}
