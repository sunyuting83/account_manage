package database

type Users struct {
	ID        uint   `gorm:"primaryKey"`
	UserName  string `gorm:"index"`
	Password  string `gorm:"index"`
	NewStatus int    `gorm:"index"`
	Remarks   string
	Projects  []Projects
	CreatedAt int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}
