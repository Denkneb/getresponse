package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	"log"
)

type ContactOpenedMessageQuery interface {
	CreateContactOpenedMessage(contactOpenedMessage *datastruct.ContactOpenedMessage) error
}

type contactOpenedMessageQuery struct{}

func (q *contactOpenedMessageQuery) CreateContactOpenedMessage(contactOpenedMessage *datastruct.ContactOpenedMessage) error {
	db := dbObj()
	err := db.Select(
		"OccurredAt",
		"Contact",
		"Account",
		"Message",
	).Create(&contactOpenedMessage).Error
	if err != nil {
		log.Println(err)
		return errors.New("connot create contactOpenedMessage")
	}
	return nil
}
