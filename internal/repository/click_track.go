package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	log "github.com/sirupsen/logrus"
)

type ClickTrackQuery interface {
	CreateClickTrack(clickTrack *datastruct.ClickTrack) error
}

type clickTrackQuery struct{}

func (q *clickTrackQuery) CreateClickTrack(clickTrack *datastruct.ClickTrack) error {
	db := dbObj()
	err := db.Where("click_track_id = ?", clickTrack.ClickTrackId).First(&clickTrack).Error
	if err != nil {
		err = db.Select(
			"ClickTrackId",
			"Name",
			"Href",
			"Url",
		).Create(&clickTrack).Error
		if err != nil {
			log.Info(err)
			return errors.New("connot create clickTrack")
		}
	}
	return nil
}
