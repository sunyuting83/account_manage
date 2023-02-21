package database

type Manager struct {
	ID             uint `gorm:"primaryKey"`
	ManageGroupsID uint
	UserName       string `gorm:"index"`
	Password       string
	NewStatus      int   `gorm:"index"`
	CreatedAt      int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt      int64 `gorm:"autoUpdateTime:milli"`
}
