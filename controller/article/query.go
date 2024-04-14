package article

import (
	"BLACKBLOG/config"
	"BLACKBLOG/controller"
	"BLACKBLOG/controller/tool"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type QueryData struct {
	Title     string `form:"title" json:"title" url:"title" default:""`                // 按标题模糊查询
	Author    string `form:"author" json:"author" url:"author" default:""`             //按作者查询
	StartTime string `form:"start_time" json:"start_time" url:"start_time" default:""` //按时间查询
	EndTime   string `form:"end_time" json:"end_time"  url:"end_time" default:""`
	//Page      int        `form:"page" json:"page" url:"page" default:"0"`
	Sort  string `form:"sort" json:"sort" url:"sort" default:""`
	Token string `form:"page_token" json:"page_token" default:""` //基于游标分页
}

type ReArticle struct {
	Id      int
	Title   string
	Content string
	Author  string
	Time    string
	Sort    string
}
type Result struct {
	Data      []ReArticle
	Respond   controller.Respond
	PageToken string `json:"page_token" default:""`
}

func Query(c *gin.Context) {
	var data QueryData
	//data.Title = c.DefaultQuery("title", "")
	//data.Author = c.DefaultQuery("author", "")
	//data.Sort = c.DefaultQuery("sort", "")
	//data.StartTime = c.DefaultQuery("start_time", "")
	//data.EndTime = c.DefaultQuery("end_time", "")
	var err error
	//data.Page, _ = strconv.Atoi(c.DefaultQuery("page", "0"))
	err = c.ShouldBindJSON(&data)
	if err != nil && err != io.EOF {
		log.SugaredLogger.Errorf("绑定数据失败:%s", err)
		c.JSON(200, controller.FailedBind)
		return
	}

	//limit, offset := tool.Pagination(data.Page)  //基于偏移量分页
	page := tool.Token(data.Token).Decode()

	if page.PageSize == 0 {
		page.PageSize = config.Conf.DataBase.Limit
	}

	var db = dao.DB
	//按照作者查询
	if data.Author != "" {
		db = db.Where("author=?", data.Author)
	}
	//按标题查询
	if data.Title != "" {
		db = db.Where("title like ?", "%"+data.Title+"%")
	}
	//按时间查询
	//判断时间范围是否合理
	if data.StartTime != "" && data.EndTime != "" && data.StartTime > data.EndTime {
		//c.JSON(200, Result{nil, controller.BadTime})
		c.JSON(200, Result{nil, controller.BadTime, ""}) //基于游标
		return
	}

	if data.StartTime != "" || data.EndTime != "" {
		var StartTime, EndTime time.Time
		var err error
		if data.StartTime != "" {
			StartTime, err = time.Parse("2006-01-02 15:04:05", data.StartTime)
			if err != nil {
				log.SugaredLogger.Errorf("字符串转换时间失败:%s", err)
			}
		}
		if data.EndTime != "" {
			EndTime, err = time.Parse("2006-01-02 15:04:05", data.EndTime)
			if err != nil {
				log.SugaredLogger.Errorf("字符串转换时间失败:%s", err)
			}

		}

		if data.StartTime != "" {
			db = db.Where("time >= ?", StartTime)
		}
		if data.EndTime != "" {
			db = db.Where("time <= ?", EndTime)
		}

	}
	//按照分类
	if data.Sort != "" {
		db = db.Where("sort=?", data.Sort)
	}
	db = db.Order("time")
	results, newCur, ok := dao.QuerysByCur(db, dao.Article{}, page.PageSize, page.NextId)
	//results, ok := dao.Querys(db, dao.Article{}, limit, offset)  //基于偏移量
	if !ok {
		//c.JSON(200, Result{nil, controller.FailedQuery})
		c.JSON(200, Result{nil, controller.FailedQuery, ""})
		return

	}

	//时间格式化
	var re = make([]ReArticle, len(results))
	for i, v := range results {
		re[i].Time = v.Time.Format("2006-01-02 15:04:05")
		re[i].Title = v.Title
		re[i].Author = v.Author
		re[i].Id = v.Id
		re[i].Content = v.Content
		re[i].Sort = v.Sort

	}
	pageToken := tool.Page{NextId: newCur, NextTimeAtuTC: time.Now().Unix(), PageSize: page.PageSize}.Encode()
	c.JSON(200, Result{re, controller.OK, string(pageToken)})
	//c.JSON(200, Result{re, controller.OK})
}
