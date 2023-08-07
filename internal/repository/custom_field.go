package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	log "github.com/sirupsen/logrus"
)

type CustomFieldQuery interface {
	CreateCustomField(customField *datastruct.CustomField) error
}

type customFieldQuery struct{}

func (q *customFieldQuery) CreateCustomField(customField *datastruct.CustomField) error {
	db := dbObj()
	err := db.Where("field_id = ?", customField.FieldId).First(&customField).Error
	if err != nil {
		err = db.Select("FieldId", "Href").Create(&customField).Error
		if err != nil {
			log.Info(err)
			return errors.New("connot create customField")
		}
	}
	return nil
}
