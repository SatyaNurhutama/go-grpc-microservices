package models

type Book struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
