package datastruct

type Message struct {
	ID           int64  `gorm:"id"`
	ResourceId   string `gorm:"resource_id"`
	ResourceType string `gorm:"resource_type"`
	Name         string `gorm:"name"`
	Href         string `gorm:"href"`
	Subject      string `gorm:"subject"`
}

func (Message) TableName() string {
	return "messages"
}
