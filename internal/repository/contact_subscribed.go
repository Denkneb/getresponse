package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	log "github.com/sirupsen/logrus"
)

type ContactSubscribedQuery interface {
	CreateContactSubscribed(contactSubscribed *datastruct.ContactSubscribed) error
}

type contactSubscribedQuery struct{}

func (q *contactSubscribedQuery) CreateContactSubscribed(contactSubscribed *datastruct.ContactSubscribed) error {
	db := dbObj()
	err := db.Select(
		"OccurredAt",
		"Contact",
		"Account",
	).Create(&contactSubscribed).Error
	if err != nil {
		log.Info(err)
		return errors.New("connot create contactSubscribed")
	}
	return nil
}
