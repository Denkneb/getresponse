package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	"log"
)

type CampaignQuery interface {
	CreateCampaign(campaign *datastruct.Campaign) error
}

type campaignQuery struct{}

func (q *campaignQuery) CreateCampaign(campaign *datastruct.Campaign) error {
	db := dbObj()
	err := db.Where("campaign_id = ?", campaign.CampaignId).First(&campaign).Error
	if err != nil {
		err = db.Select("CampaignId", "Name", "Href").Create(&campaign).Error
		if err != nil {
			log.Println(err)
			return  errors.New("connot create campaign")
		}
	}
	return nil
}
