package repositories

import (
	"database/sql"

	"githup.com/Therocking/go-http/db"
	"githup.com/Therocking/go-http/models"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := db.Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.FullName, &user.LastName, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(userId int64) (models.User, error) {
	query := "SELECT * FROM users WHERE id = ?"

	var user models.User
	err := db.Db.QueryRow(query, userId).Scan(&user.Id, &user.FullName, &user.LastName, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CreateUser(user models.User) (sql.Result, error) {
	query := "INSERT INTO users(fullname, lastname, username, password) VALUES (?, ?, ?, ?)"
	result, err := db.Db.Exec(query, user.FullName, user.LastName, user.Username, user.Password)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateUser(user models.User) (models.User, error) {
	query := "UPDATE users SET fullname = ?, username = ?, password = ? WHERE id = ?"

	err := db.Db.QueryRow(query, user.FullName, user.Username, user.Password, user.Id).Scan(
		&user.Id, &user.FullName, &user.LastName, &user.Username, &user.Password)

	if err != nil {
		return models.User{}, err

	}
	return user, nil
}

func DeleteUser(userId int64) (models.User, error) {
	query := "DELETE FROM users WHERE id = ?"

	var user models.User
	err := db.Db.QueryRow(query, userId).Scan(&user.Id, &user.FullName, &user.LastName, &user.Username, &user.Password)

	if err != nil {
		return models.User{}, err

	}

	return user, nil
}
