package services

import (
	"database/sql"
	"fmt"

	"githup.com/Therocking/go-http/models"
	"githup.com/Therocking/go-http/repositories"
)

var usersList []models.User = []models.User{}

func GetAllUsers() ([]models.User, error) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}

func GetUser(userId int64) (models.User, error) {
	user, err := repositories.GetUser(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("models.User with id: %v was not found", userId)
		}
		return models.User{}, err
	}

	return user, nil
}

func CreateUser(user models.User) (string, error) {
	result, err := repositories.CreateUser(user)
	if err != nil {
		return "", err
	}
	id, _ := result.LastInsertId()

	return fmt.Sprintf("models.User with id: %v was added", id), nil
}

func UpdateUser(user models.User) (models.User, error) {
	resp, err := repositories.UpdateUser(user)

	if err != nil {
		return models.User{}, err
	}

	return resp, nil
}

func DeleteUser(userId int64) (models.User, error) {
	user, err := repositories.DeleteUser(userId)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
