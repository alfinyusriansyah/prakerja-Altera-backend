package models

type Book struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Year      int    `json:"year"`
	Author    string `json:"author"`
	PageCount int    `json:"pagecount"`
}
