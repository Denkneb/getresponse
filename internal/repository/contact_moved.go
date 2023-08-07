package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	log "github.com/sirupsen/logrus"
)

type ContactMovedQuery interface {
	CreateContactMoved(contactMoved *datastruct.ContactMoved) error
}

type contactMovedQuery struct{}

func (q *contactMovedQuery) CreateContactMoved(contactMoved *datastruct.ContactMoved) error {
	db := dbObj()
	err := db.Select(
		"OccurredAt",
		"Contact",
		"Account",
		"SourceCampaign",
		"CampaignTarget",
	).Create(&contactMoved).Error
	if err != nil {
		log.Info(err)
		return errors.New("connot create contactMoved")
	}
	return nil
}
