package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	"log"
)

type ContactUnsubscribedQuery interface {
	CreateContactUnsubscribed(contactUnsubscribed *datastruct.ContactUnsubscribed) error
}

type contactUnsubscribedQuery struct{}

func (q *contactUnsubscribedQuery) CreateContactUnsubscribed(contactUnsubscribed *datastruct.ContactUnsubscribed) error {
	db := dbObj()
	err := db.Select(
		"OccurredAt",
		"Contact",
		"Account",
	).Create(&contactUnsubscribed).Error
	if err != nil {
		log.Println(err)
		return errors.New("connot create contactUnsubscribed")
	}
	return nil
}
