package tool

import "BLACKBLOG/config"

func Pagination(page int) (limit, offset int) {
	limit = config.Conf.DataBase.Limit
	offset = (page - 1) * limit
	return
}
