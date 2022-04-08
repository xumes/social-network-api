package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Posts) (uint64, error) {
	sql := "INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)"

	statement, err := repository.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorId)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}
