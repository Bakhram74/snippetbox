package mysql

import (
	"database/sql"
	"errors"
	"github.com/Bakhram74/snippetbox.git/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title,content,created,expires)
VALUES(?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), nil
}
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title,content,created,expires from snippets
WHERE expires > UTC_TIMESTAMP() AND id = ?`
	s := &models.Snippet{}
	err := m.DB.QueryRow(stmt, id).Scan(&s.Id, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}
func (m *SnippetModel) Latest(id int) ([]models.Snippet, error) {
	return nil, nil
}
