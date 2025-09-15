package repository

import (
	"database/sql"
	"wiki/backend/internal/model"
)

func GetPageBySlug(db *sql.DB, slug string) (*model.Page, error) {
	page := &model.Page{}
	err := db.QueryRow("SELECT id, title, slug, content, last_editor_id FROM pages WHERE slug = ?", slug).
		Scan(&page.ID, &page.Title, &page.Slug, &page.Content, &page.LastEditorID)
	if err != nil {
		return nil, err
	}
	return page, nil
}
func CreatePage(db *sql.DB, title, slug, content string, lastEditorID int) (int64, error) {
	result, err := db.Exec("INSERT INTO pages (title, slug, content, last_editor_id) VALUES (?, ?, ?, ?)", title, slug, content, lastEditorID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
func UpdatePageContent(db *sql.DB, id int, content string, lastEditorID int) error {
	_, err := db.Exec("UPDATE pages SET content = ?, last_editor_id = ? WHERE id = ?", content, lastEditorID, id)
	return err
}
func DeletePage(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM pages WHERE id = ?", id)
	return err
}
