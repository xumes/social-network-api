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

func (repository Posts) GetByid(id uint64) (models.Posts, error) {
	sql := `SELECT posts.id, posts.title, posts.content, posts.author_id, posts.likes, posts.created_at, users.nick
				FROM posts INNER JOIN users
				ON users.id = posts.author_id
				WHERE posts.id = ?
				`
	row, err := repository.db.Query(sql, id)
	if err != nil {
		return models.Posts{}, err
	}
	defer row.Close()

	var post models.Posts

	if row.Next() {
		if err = row.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Posts{}, err
		}
	}

	return post, nil
}

func (repository Posts) GetAllFromUser(userId uint64) ([]models.Posts, error) {
	sql := `SELECT DISTINCT
					posts.id, posts.title, posts.content, posts.author_id, posts.likes, posts.created_at, users.nick
				FROM
					posts
				INNER JOIN
					users
					ON users.id = posts.author_id
				INNER JOIN
					followers
					ON posts.author_id = followers.user_id
				WHERE
					users.id = ?
					OR followers.follower_id = ?
				ORDER BY
					posts.id DESC
				`
	rows, err := repository.db.Query(sql, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Posts

	for rows.Next() {
		var post models.Posts

		if err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) GetPostAuthorId(postId uint64) (uint64, error) {
	sql := `SELECT posts.author_id
				FROM posts
				WHERE posts.id = ?
				`
	row, err := repository.db.Query(sql, postId)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var authorId uint64

	if row.Next() {
		if err = row.Scan(
			&authorId,
		); err != nil {
			return 0, err
		}
	}

	return authorId, nil
}

func (repository Posts) UpdateById(id uint64, post models.Posts) error {
	sql := "UPDATE posts SET title = ?, content = ? WHERE id = ?"

	statement, err := repository.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, id); err != nil {
		return err
	}

	return nil
}

func (repository Posts) RemoveById(id uint64) error {
	sql := "DELETE FROM posts WHERE id = ?"

	statement, err := repository.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (repository Posts) GetByUserId(userId uint64) ([]models.Posts, error) {
	sql := `SELECT posts.id, posts.title, posts.content, posts.author_id, posts.likes, posts.created_at, users.nick
				FROM posts INNER JOIN users
				ON users.id = posts.author_id
				WHERE posts.author_id = ?
				`
	row, err := repository.db.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var posts []models.Posts

	for row.Next() {
		var post models.Posts

		if err = row.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) Like(postId uint64) error {
	sql := "UPDATE posts SET likes = likes + 1 WHERE id = ?"

	statement, err := repository.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (repository Posts) Unlike(postId uint64) error {
	sql := "UPDATE posts SET likes = likes - 1 WHERE likes > 0 AND id = ?"

	statement, err := repository.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}
