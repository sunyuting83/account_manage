package database

type Manager struct {
	ID        uint   `gorm:"primaryKey"`
	UserName  string `gorm:"index"`
	Password  string
	NewStatus int   `gorm:"index"`
	CreatedAt int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}

func CheckAdminLogin(username, password string) (manager *Manager, err error) {
	if err = sqlDB.First(&manager, "user_name = ? AND new_status = ? AND password = ?", username, "0", password).Error; err != nil {
		return
	}
	return
}
