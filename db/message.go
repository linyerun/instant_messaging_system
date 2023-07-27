package db

import "instant_messaging_system/models"

func SaveMessage(msg *models.Message) error {
	return MySQLClient.Create(msg).Error
}
