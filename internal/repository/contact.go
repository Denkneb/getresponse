package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	log "github.com/sirupsen/logrus"
)

type ContactQuery interface {
	CreateContact(contact *datastruct.Contact) error
}

type contactQuery struct{}

func (q *contactQuery) CreateContact(contact *datastruct.Contact) error {
	db := dbObj()
	err := db.Where("contact_id = ?", contact.ContactId).First(&contact).Error
	fields := []string{
		"ContactId",
		"Email",
		"Name",
		"Ip",
		"Origin",
		"Href",
		"Campaign",
	}
	if contact.PhoneNumber != 0 {
		fields = append(fields, "PhoneNumber")
	}
	if err != nil {
		err = db.Select(fields).Create(&contact).Error
		if err != nil {
			log.Info(err)
			return errors.New("connot create contact")
		}
	}
	return nil
}
