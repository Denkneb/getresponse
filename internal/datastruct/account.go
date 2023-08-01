package datastruct

type Account struct {
	ID        int64  `gorm:"id"`
	AccountId string `gorm:"account_id"`
}

func (Account) TableName() string {
	return "accounts"
}
