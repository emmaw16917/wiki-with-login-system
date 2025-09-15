package service

import (
	"database/sql"
	"errors"
	"wiki/backend/internal/model"
	"wiki/backend/internal/repository"
)

func GetPageBySlug(db *sql.DB, slug string) (*model.Page, error) {
	return repository.GetPageBySlug(db, slug)
}

func CreatePage(db *sql.DB, page *model.Page, userID int) (int64, error) {
	if userID == 0 {
		return 0, errors.New("未登录用户不能新建页面")
	}
	page.LastEditorID = userID
	return repository.CreatePage(db, page.Title, page.Slug, page.Content, page.LastEditorID)
}

func UpdatePageContent(db *sql.DB, id int, content string, userID int) error {
	if userID == 0 {
		return errors.New("未登录用户不能编辑页面")
	}
	return repository.UpdatePageContent(db, id, content, userID)
}

func DeletePage(db *sql.DB, id int, userID int) error {
	if userID == 0 {
		return errors.New("未登录用户不能删除页面")
	}
	return repository.DeletePage(db, id)
}
