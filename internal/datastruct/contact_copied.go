package datastruct

import "time"

type ContactCopied struct {
	ID               int64     `gorm:"id"`
	OccurredAt       time.Time `gorm:"occurred_at"`
	IsSentToKis      bool      `gorm:"is_sent_to_kis"`
	SentToKisAt      time.Time `gorm:"sent_to_kis_at"`
	IsDeliveredToKis bool      `gorm:"is_delivered_to_kis"`
	SentLocation     string    `gorm:"sent_location"`
	Contact          int64     `gorm:"contact"`
	Account          int64     `gorm:"account"`
	SourceCampaign   int64     `gorm:"source_campaign"`
	CampaignTarget   int64     `gorm:"campaign_target"`
}

func (ContactCopied) TableName() string {
	return "contacts_copied"
}
