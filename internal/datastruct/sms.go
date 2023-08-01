package datastruct

type SMS struct {
	ID    int64  `gorm:"id"`
	SmsId string `gorm:"sms_id"`
	Name  string `gorm:"name"`
	Href  string `gorm:"href"`
}

func (SMS) TableName() string {
	return "sms"
}
