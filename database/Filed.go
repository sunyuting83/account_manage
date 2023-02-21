package database

type Filed struct {
	ID         uint `gorm:"primaryKey"`
	ProjectsID uint
	FiledName  string
	Data       string
	CreatedAt  int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt  int64 `gorm:"autoUpdateTime:milli"`
}
