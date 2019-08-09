package services

import (
	"../database"
	"../models"
)

func UserRegister(name string, age int, height int) error {
	db := database.DB
	trx := db.Begin()
	if err := trx.Create(&models.User{Name: name, Age: age, Height: height}).Error; err != nil {
		trx.Rollback()
		return err
	}
	trx.Commit()
	return nil
}
func GetUserByID(id int) (*models.User, error) {
	db := database.DB
	// user := &models.User{ID: id}
	user := &models.User{}
	res := db.First(&user, id)
	err := res.Error
	return user, err
}
func GetUserByName(name string) []*models.User {
	db := database.DB
	var users []*models.User
	db.Where(&models.User{Name: name}).Find(&users)
	return users
}

func GetUserByToken(token string) *models.User {
	db := database.DB
	user := &models.User{}
	tokenObj := &models.Token{}
	res := db.Where(&models.Token{Token: token, Status: "E"}).First(&tokenObj)
	if res.Error != nil {
		return nil
	}
	db.First(&user, tokenObj.UserID)
	return user
}

// func GetUserByName(name string) (users []*models.User) {
// 	db := database.DB
// 	db.Where(&models.User{Name: "gbs"}).Find(&users)
// 	return
// }
