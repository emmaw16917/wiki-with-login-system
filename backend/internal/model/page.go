package model

type Page struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Content      string `json:"content"`
	LastEditorID int    `json:"last_editor_id"`
}
