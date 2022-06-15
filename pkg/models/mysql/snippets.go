package mysql

import (
	"database/sql"
	"github.com/Bakhram74/snippetbox.git/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `insert into snippets (title,content,created,expires)
values(?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), nil
}
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}
func (m *SnippetModel) Latest(id int) ([]models.Snippet, error) {
	return nil, nil
}
