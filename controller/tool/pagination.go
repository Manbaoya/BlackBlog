package tool

import (
	"BLACKBLOG/config"
	"BLACKBLOG/log"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func Pagination(page int) (limit, offset int) {
	limit = config.Conf.DataBase.Limit
	offset = (page - 1) * limit
	return
}

type Page struct {
	NextId        int   `json:"next_id"`          //游标
	NextTimeAtuTC int64 `json:"next_time_at_utc"` //分页
	PageSize      int   `json:"page_size"`        //每页的元素个数
}
type Token string

func (p Page) Encode() Token { //返回分页token
	b, err := json.Marshal(p)
	if err != nil {
		log.SugaredLogger.Error("返回分页失败: ", err)
		return Token("")
	}
	return Token(base64.StdEncoding.EncodeToString(b))
}

func (token Token) Decode() Page { //解析分页token
	var result Page
	if len(token) == 0 {
		return result
	}
	bytes, err := base64.StdEncoding.DecodeString(string(token))
	if err != nil {
		log.SugaredLogger.Error("解析分页失败：", err)
		return result
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		log.SugaredLogger.Error("解析分页失败：", err)
		return result
	}
	fmt.Println(result)
	return result

}
