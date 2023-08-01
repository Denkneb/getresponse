package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DAO interface {
	NewAccountQuery() AccountQuery
	NewCampaignQuery() CampaignQuery
	NewContactQuery() ContactQuery
	NewCustomFieldQuery() CustomFieldQuery
	NewContactSubscribedQuery() ContactSubscribedQuery
	NewContactUnsubscribedQuery() ContactUnsubscribedQuery
	NewContactCopiedQuery() ContactCopiedQuery
	NewContactMovedQuery() ContactMovedQuery
	NewContactOpenedMessageQuery() ContactOpenedMessageQuery
	NewContactLinkClickedQuery() ContactLinkClickedQuery
	NewContactSmsLinkClickedQuery() ContactSmsLinkClickedQuery
	NewMessageQuery() MessageQuery
	NewClickTrackQuery() ClickTrackQuery
	NewSMSQuery() SMSQuery
}

type dao struct{}

var DB *sql.DB

func dbObj() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: DB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot initialized db session: %v", err)
	}
	return db
}

func NewDAO(db *sql.DB) DAO {
	DB = db
	return &dao{}
}

func NewDB() (*sql.DB, error) {
	host := viper.Get("database.host").(string)
	port := viper.Get("database.port").(int)
	user := viper.Get("database.user").(string)
	dbname := viper.Get("database.dbname").(string)
	password := viper.Get("database.password").(string)

	// Starting a database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d *dao) NewAccountQuery() AccountQuery {
	return &accountQuery{}
}

func (d *dao) NewCampaignQuery() CampaignQuery {
	return &campaignQuery{}
}

func (d *dao) NewContactQuery() ContactQuery {
	return &contactQuery{}
}

func (d *dao) NewCustomFieldQuery() CustomFieldQuery {
	return &customFieldQuery{}
}

func (d *dao) NewContactSubscribedQuery() ContactSubscribedQuery {
	return &contactSubscribedQuery{}
}

func (d *dao) NewContactUnsubscribedQuery() ContactUnsubscribedQuery {
	return &contactUnsubscribedQuery{}
}

func (d *dao) NewContactCopiedQuery() ContactCopiedQuery {
	return &contactCopiedQuery{}
}

func (d *dao) NewContactMovedQuery() ContactMovedQuery {
	return &contactMovedQuery{}
}

func (d *dao) NewContactOpenedMessageQuery() ContactOpenedMessageQuery {
	return &contactOpenedMessageQuery{}
}

func (d *dao) NewContactLinkClickedQuery() ContactLinkClickedQuery {
	return &contactLinkClickedQuery{}
}

func (d *dao) NewContactSmsLinkClickedQuery() ContactSmsLinkClickedQuery {
	return &contactSmsLinkClickedQuery{}
}

func (d *dao) NewMessageQuery() MessageQuery {
	return &messageQuery{}
}

func (d *dao) NewClickTrackQuery() ClickTrackQuery {
	return &clickTrackQuery{}
}

func (d *dao) NewSMSQuery() SMSQuery {
	return &smsQuery{}
}
