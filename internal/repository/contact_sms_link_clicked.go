package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	"log"
)

type ContactSmsLinkClickedQuery interface {
	CreateContactSmsLinkClicked(contactSmsLinkClicked *datastruct.ContactSmsLinkClicked) error
}

type contactSmsLinkClickedQuery struct{}

func (q *contactSmsLinkClickedQuery) CreateContactSmsLinkClicked(contactSmsLinkClicked *datastruct.ContactSmsLinkClicked) error {
	db := dbObj()
	err := db.Select(
		"OccurredAt",
		"Contact",
		"Account",
		"SMS",
		"ClickTrack",
	).Create(&contactSmsLinkClicked).Error
	if err != nil {
		log.Println(err)
		return errors.New("connot create contactSmsLinkClicked")
	}
	return nil
}
