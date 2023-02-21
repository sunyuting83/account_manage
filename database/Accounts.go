package database

type Accounts struct {
	ID            uint `gorm:"primaryKey"`
	ProjectsID    uint
	UserName      string
	Password      string
	Cover         string
	NewStatus     int `gorm:"index"`
	TodayGold     int64
	YesterdayGold int64
	Multiple      int64
	Diamond       int
	Crazy         int
	Precise       int
	Cold          int
	Exptime       int64
	Remarks       string
	CreatedAt     int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt     int64 `gorm:"autoUpdateTime:milli"`
}
