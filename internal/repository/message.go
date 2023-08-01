package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	"log"
)

type MessageQuery interface {
	CreateMessage(message *datastruct.Message) error
}

type messageQuery struct{}

func (q *messageQuery) CreateMessage(message *datastruct.Message) error {
	db := dbObj()
	err := db.Where("resource_id = ?", message.ResourceId).First(&message).Error
	if err != nil {
		err = db.Select(
			"ResourceId",
			"ResourceType",
			"Name",
			"Href",
			"Subject",
		).Create(&message).Error
		if err != nil {
			log.Println(err)
			return errors.New("connot create message")
		}
	}
	return nil
}
