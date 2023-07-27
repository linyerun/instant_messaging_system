package db

import "instant_messaging_system/models"

func GetUserCountByEmail(email string) (count int64) {
	MySQLClient.Model(models.User{}).Where("email = ?", email).Count(&count)
	return
}

func InsertUser(user *models.User) error {
	return MySQLClient.Create(user).Error
}

func GetUserIdByEmailPwd(email, pwd string) (id uint, err error) {
	err = MySQLClient.Model(models.User{}).Select("id").Where("email = ? and password = ?", email, pwd).Row().Scan(&id)
	return
}

func GetUserById(uid uint) (user *models.User, err error) {
	user = new(models.User)
	err = MySQLClient.First(user, "id = ?", uid).Error
	return
}

func GetUserByEmail(email string) (user *models.User, err error) {
	user = new(models.User)
	err = MySQLClient.First(user, "email = ?", email).Error
	return
}
