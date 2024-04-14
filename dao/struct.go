package dao

import "time"

type T interface {
	User | Article | Image | Comment
}

//	type S interface {
//		[]Article | []Comment
//	}
type User struct {
	Id       int    `gorm:"id"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Phone    string `gorm:"phone"`
}
type Article struct {
	Id      int       `gorm:"id"`
	Title   string    `gorm:"title"`
	Content string    `gorm:"content"`
	Author  string    `gorm:"author"`
	Time    time.Time `gorm:"time"`
	Sort    string    `gorm:"sort"`
}
type Image struct {
	Id   int    `gorm:"id"`
	Name string `gorm:"name"`
	Path string `gorm:"path"`
}

type Comment struct {
	Id         int       `gorm:"id"`
	ArticleId  int       `gorm:"article_id"`
	UserId     int       `gorm:"user_id"`
	Content    string    `gorm:"content"`
	CreateTime time.Time `gorm:"create_time"`
}
