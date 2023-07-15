package models

type Book struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Year      int    `json:"year"`
	PageCount int    `json:"page_count" gorm:"not null"`
	AuthorId  int    `json:"author_id"`
}
