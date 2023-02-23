package database

type ManageGroups struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Power     string
	Manager   []Manager
	CreatedAt int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}
