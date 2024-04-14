package user

import (
	"BLACKBLOG/controller"
	"BLACKBLOG/dao"
	"BLACKBLOG/log"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type ImageResponse struct {
	Respond  controller.Respond
	ImageUrl string `json:"imageUrl"`
}

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.SugaredLogger.Errorf("上传文件失败:%s", err)
		c.JSON(200, controller.FailedUpload)
		return
	}
	i, exist := c.Get("id")
	id := i.(int)
	if !exist {
		c.JSON(200, controller.FailedGetId)
		return
	}
	ExtName := path.Ext(file.Filename)
	allow := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
	}
	if _, ok := allow[ExtName]; !ok {
		c.JSON(200, controller.BadFile)
		return
	}
	currentTime := time.Now().Format("20060102")
	if err := os.MkdirAll("upload/upload/"+currentTime, 0755); err != nil {
		log.SugaredLogger.Errorf("创建文件目录失败:%s", err)
		c.JSON(200, controller.FailedUpload)
		return
	}
	fileUnixName := strconv.FormatInt(time.Now().UnixNano(), 10)
	saveDir := path.Join("upload/upload/"+currentTime, fileUnixName)
	err = c.SaveUploadedFile(file, saveDir)
	if err != nil {
		log.SugaredLogger.Errorf("文件保存失败:%s", err)
		c.JSON(200, controller.FailedSaved)
		return
	}
	imageUrl := strings.Replace(saveDir, "upload/upload", "http://localhost:8080/image", -1)
	db := dao.DB
	db.Where("user_id=?", id)
	if _, ok := dao.Query(db, dao.Image{Id: id}); !ok {
		if _, k := dao.Add(dao.DB, dao.Image{Id: id, Name: fileUnixName, Path: saveDir}); !k {
			c.JSON(200, ImageResponse{controller.FailedCreate, imageUrl})
			return
		}
	} else {
		if k := dao.Alter(dao.DB, id, dao.Image{Id: id, Name: fileUnixName, Path: saveDir}); !k {
			c.JSON(200, ImageResponse{controller.FailedAlter, imageUrl})
			return
		}
	}
	c.JSON(200, ImageResponse{controller.OK, imageUrl})

}
