package database

type Accounts struct {
	ID            uint `gorm:"primaryKey"`
	ProjectsID    uint
	ComputID      uint
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

// Get Count
func (accounts *Accounts) GetCount(ProjectsID string) (count int64, err error) {
	if err = sqlDB.Model(&accounts).Where("projects_id = ?", ProjectsID).Count(&count).Error; err != nil {
		return
	}
	return
}

// Account List
func GetAccountList(page int) (accounts *[]Accounts, err error) {
	p := makePage(page)
	if err = sqlDB.
		Order("updated_at desc").
		Limit(100).Offset(p).
		Find(&accounts).Error; err != nil {
		return
	}
	return
}
