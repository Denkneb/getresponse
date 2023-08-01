package datastruct

type Campaign struct {
	ID         int64  `gorm:"id"`
	CampaignId string `gorm:"campaign_id"`
	Name       string `gorm:"name"`
	Href       string `gorm:"href"`
}

func (Campaign) TableName() string {
	return "campaigns"
}
