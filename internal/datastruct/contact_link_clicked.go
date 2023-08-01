package datastruct

import "time"

type ContactLinkClicked struct {
	ID               int64     `gorm:"id"`
	OccurredAt       time.Time `gorm:"occurred_at"`
	IsSentToKis      bool      `gorm:"is_sent_to_kis"`
	SentToKisAt      time.Time `gorm:"sent_to_kis_at"`
	IsDeliveredToKis bool      `gorm:"is_delivered_to_kis"`
	SentLocation     string    `gorm:"sent_location"`
	Contact          int64     `gorm:"contact"`
	Account          int64     `gorm:"account"`
	Message          int64     `gorm:"message"`
	ClickTrack       int64     `gorm:"click_track"`
}

func (ContactLinkClicked) TableName() string {
	return "links_clicked"
}
