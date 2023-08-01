package repository

import (
	"errors"
	"getresponse/internal/datastruct"
	"log"
)

type AccountQuery interface {
	CreateAccount(account *datastruct.Account) error
}

type accountQuery struct{}

func (q *accountQuery) CreateAccount(account *datastruct.Account) error {
	db := dbObj()
	err := db.Where("account_id = ?", account.AccountId).First(&account).Error
	if err != nil {
		err = db.Select("AccountId").Create(&account).Error
		if err != nil {
			log.Println(err)
			return errors.New("connot create account")
		}
	}
	return nil
}
