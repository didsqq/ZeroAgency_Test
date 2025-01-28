package models

type News struct {
	Id         int64   `json:"Id" reform:"id"`
	Title      string  `json:"Title" reform:"title"`
	Content    string  `json:"Content" reform:"content"`
	Categories []int64 `json:"Categories" reform:"categories"`
}
