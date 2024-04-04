package services

import (

	"githup.com/Therocking/go-http/models"
	"githup.com/Therocking/go-http/repositories"
)

func GetAllUsers() ([]models.User, error) {
	users := repositories.GetAllUsers()

	return users, nil
}

func GetUser(userId int64) (models.User, error) {
	user, err := repositories.GetUser(userId)
	if err != nil {
		
		return models.User{}, err
	}

	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	user, err := repositories.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdateUser(id int64, user models.User) (models.User, error) {
	user, err := repositories.UpdateUser(id, user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func DeleteUser(userId int64) (models.User, error) {
	user, err := repositories.DeleteUser(userId)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
