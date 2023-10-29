package repositories

import (
	"notepad-api/configs"
	"notepad-api/models"
)

func EmailChecking(existingUser *models.User, user *models.User) error {
	err := configs.DB.Where("email = ?", user.Email).First(&existingUser).Error
	return err
}

func AddUser(user *models.User) error {
	err := configs.DB.Create(&user).Error
	return err
}

func GetOneUser(idUser uint, oneuser *models.User) error {
	err := configs.DB.First(&oneuser, idUser).Error
	return err
}

func DeleteOneUser(idUser uint, oneuser *models.User) error {
	err := configs.DB.Delete(&oneuser, idUser).Error
	return err
}
