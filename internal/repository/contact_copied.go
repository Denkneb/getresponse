package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	"log"
)

type ContactCopiedQuery interface {
	CreateContactCopied(contactCopied *datastruct.ContactCopied) error
}

type contactCopiedQuery struct{}

func (q *contactCopiedQuery) CreateContactCopied(contactCopied *datastruct.ContactCopied) error {
	db := dbObj()
	err := db.Select(
		"OccurredAt",
		"Contact",
		"Account",
		"SourceCampaign",
		"CampaignTarget",
	).Create(&contactCopied).Error
	if err != nil {
		log.Println(err)
		return errors.New("connot create contactCopied")
	}
	return nil
}
