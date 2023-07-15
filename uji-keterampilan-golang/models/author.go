package models

type Author struct {
	Id    int    `json:"id" gorm:"primaryKey" validate:"required"`
	Name  string `json:"name" gorm:"not null"`
	Books []Book `json:"books"`
}
