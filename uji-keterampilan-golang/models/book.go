package models

import "time"

type Book struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Year      int       `json:"year"`
	Author    string    `json:"author"`
	PageCount int       `json:"pagecount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
