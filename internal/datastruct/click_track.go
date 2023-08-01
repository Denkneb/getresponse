package datastruct

type ClickTrack struct {
	ID           int64  `gorm:"id"`
	ClickTrackId string `gorm:"click_track_id"`
	Name         string `gorm:"name"`
	Href         string `gorm:"href"`
	Url          string `gorm:"url"`
}

func (ClickTrack) TableName() string {
	return "clicks_track"
}
