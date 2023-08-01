package datastruct

type CustomField struct {
	ID      int64  `gorm:"id"`
	FieldId string `gorm:"field_id"`
	Href    string `gorm:"href"`
}

func (CustomField) TableName() string {
	return "custom_fields"
}
