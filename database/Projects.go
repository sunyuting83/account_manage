package database

type Projects struct {
	ID           uint `gorm:"primaryKey"`
	UsersID      uint
	ProjectsName string
	NewStatus    int `gorm:"index"`
	Accounts     []Accounts
	Filed        []Filed
	CreatedAt    int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt    int64 `gorm:"autoUpdateTime:milli"`
}
