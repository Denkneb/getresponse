package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	log "github.com/sirupsen/logrus"
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
		log.Info(err)
		return errors.New("connot create contactOpenedMessage")
	}
	return nil
}
