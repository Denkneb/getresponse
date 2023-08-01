package datastruct

import "time"

type Contact struct {
	ID          int64  `gorm:"id"`
	ContactId   string `gorm:"contactId"`
	Email       string `gorm:"email"`
	Name        string `gorm:"name"`
	Ip          string `gorm:"ip"`
	Origin      string `gorm:"origin"`
	Href        string `gorm:"href"`
	Campaign    int64  `gorm:"campaign"`
	PhoneNumber int64  `gorm:"phone_number"`
	CreatedAt   time.Time  `gorm:"created_at"`
	UpdatedAt   time.Time  `gorm:"updated_at"`
}

func (Contact) TableName() string {
	return "contacts"
}
