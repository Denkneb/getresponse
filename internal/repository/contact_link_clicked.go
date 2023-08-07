package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	log "github.com/sirupsen/logrus"
)

type ContactLinkClickedQuery interface {
	CreateContactLinkClicked(contactLinkClicked *datastruct.ContactLinkClicked) error
}

type contactLinkClickedQuery struct{}

func (q *contactLinkClickedQuery) CreateContactLinkClicked(contactLinkClicked *datastruct.ContactLinkClicked) error {
	db := dbObj()
	err := db.Select(
		"OccurredAt",
		"Contact",
		"Account",
		"Message",
		"ClickTrack",
	).Create(&contactLinkClicked).Error
	if err != nil {
		log.Info(err)
		return errors.New("connot create contactLinkClicked")
	}
	return nil
}
