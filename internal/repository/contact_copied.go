package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	log "github.com/sirupsen/logrus"
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
		log.Info(err)
		return errors.New("connot create contactCopied")
	}
	return nil
}
