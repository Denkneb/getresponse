package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	"log"
)

type SMSQuery interface {
	CreateSMS(message *datastruct.SMS) error
}

type smsQuery struct{}

func (q *smsQuery) CreateSMS(sms *datastruct.SMS) error {
	db := dbObj()
	err := db.Where("sms_id = ?", sms.SmsId).First(&sms).Error
	if err != nil {
		err = db.Select(
			"SmsId",
			"Name",
			"Href",
		).Create(&sms).Error
		if err != nil {
			log.Println(err)
			return errors.New("connot create sms")
		}
	}
	return nil
}
