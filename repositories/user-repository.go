package repositories

import (
	"githup.com/Therocking/go-http/db"
	"githup.com/Therocking/go-http/models"
)

func GetAllUsers(filters models.Filters) []models.User {
	var users []models.User
	db.Db.Limit(filters.Limit).Offset(filters.Skip).Find(&users)

	return users
}

func GetUser(userId int64) (models.User, error) {
	var user models.User

	userFound := db.Db.First(&user, userId)
	if err := userFound.Error; err != nil {
		return models.User{}, err
	}

	db.Db.Model(&user).Association("Products").Find(&user.Products)

	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	createdUser := db.Db.Create(&user)

	if err := createdUser.Error; err != nil {
		return models.User{}, err
	}

	db.Db.Model(&user).Association("Products").Find(&user.Products)

	return user, nil
}

func UpdateUser(id int64, user models.User) (models.User, error) {
	userFound := db.Db.First(models.User{}, id)
	if err := userFound.Error; err != nil {
		return models.User{}, err
	}

	user.Username = "updated"

	db.Db.Model(&user).Association("Products").Find(&user.Products)
	db.Db.Save(&user)

	return user, nil
}

func DeleteUser(userId int64) (models.User, error) {
	var user models.User
	userFound := db.Db.First(&user, userId)
	if err := userFound.Error; err != nil {
		return models.User{}, err
	}

	db.Db.Model(&user).Association("Products").Find(&user.Products)

	db.Db.Delete(&user)
	return user, nil
}
