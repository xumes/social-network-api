package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (userRepository users) Create(user models.User) (uint64, error) {
	sql := "INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)"
	statement, err := userRepository.db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}

func (userRepository users) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	sql := "SELECT id, name, nick, email, created_at FROM users WHERE  name LIKE ? OR nick LIKE  ?"
	rows, err := userRepository.db.Query(sql, nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (userRepository users) FindById(id uint64) (models.User, error) {
	sql := "SELECT  name, nick, email, created_at FROM users WHERE  id = ?"
	rows, err := userRepository.db.Query(sql, id)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
		user.Id = id
	}

	return user, nil
}

func (userRepository users) FindByEmail(email string) (models.User, error) {
	sql := "SELECT  id, password FROM users WHERE  email = ?"
	row, err := userRepository.db.Query(sql, email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(&user.Id, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (userRepository users) UpdateById(id uint64, user models.User) error {
	sql := "UPDATE users SET  name=?, nick=?, email=? WHERE id=?"
	statement, err := userRepository.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (userRepository users) DeleteById(id uint64) error {
	sql := "DELETE FROM users WHERE  id = ?"
	statement, err := userRepository.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (userRepository users) Follow(userId, followerId uint64) error {
	sql := "INSERT ignore INTO followers (user_id, follower_id) values (?, ?)"
	statement, err := userRepository.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (userRepository users) Unfollow(userId, followerId uint64) error {
	sql := "DELETE FROM followers WHERE user_id = ? AND  follower_id =  ?"
	statement, err := userRepository.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (userRepository users) Followers(userId uint64) ([]models.User, error) {
	sql := `SELECT users.id, users.name, users.nick, users.email, users.created_at
				FROM users 
				INNER JOIN followers ON  users.id = followers.follower_id
				WHERE followers.user_id = ?`
	rows, err := userRepository.db.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		followers = append(followers, user)
	}

	return followers, nil
}

func (userRepository users) Following(userId uint64) ([]models.User, error) {
	sql := `SELECT users.id, users.name, users.nick, users.email, users.created_at
				FROM users 
				INNER JOIN followers ON  users.id = followers.follower_id
				WHERE followers.follower_id = ?`
	rows, err := userRepository.db.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		followers = append(followers, user)
	}

	return followers, nil
}
