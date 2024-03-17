package dao

import "time"

type T interface {
	User | Article | Image
}
type S interface {
	[]Article
}
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
