package database

import "fmt"

type Manager struct {
	ID        uint   `gorm:"primaryKey"`
	UserName  string `gorm:"index"`
	Password  string
	NewStatus int   `gorm:"index"`
	CreatedAt int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}

func (manager *Manager) Insert() (err error) {
	sqlDB.Create(&manager)
	return nil
}

func CheckAdminLogin(username, password string) (manager *Manager, err error) {
	if err = sqlDB.First(&manager, "user_name = ? AND new_status = ? AND password = ?", username, "0", password).Error; err != nil {
		return
	}
	return
}

func CheckUserName(username string) (manager *Manager, err error) {
	if err = sqlDB.First(&manager, "user_name = ? ", username).Error; err != nil {
		return
	}
	return
}

// Get Count
func (manager *Manager) GetCount() (count int64, err error) {
	if err = sqlDB.Model(&manager).Count(&count).Error; err != nil {
		return
	}
	return
}

// Check ID
func CheckID(id int64) (manager *Manager, err error) {
	if err = sqlDB.First(&manager, "id = ?", id).Error; err != nil {
		return
	}
	return
}

// Delete Admin
func (manager *Manager) DeleteOne(id int64) {
	// time.Sleep(time.Duration(100) * time.Millisecond)
	sqlDB.Where("id = ?", id).Delete(&manager)
}

// Admin List
func GetAdminList(page int) (manages *[]Manager, err error) {
	p := makePage(page)
	if err = sqlDB.
		Select("id, user_name, new_status, created_at").
		Order("id desc").
		Limit(100).Offset(p).
		Find(&manages).Error; err != nil {
		return
	}
	return
}

// Reset Password
func (manager *Manager) ResetPassword(username string) (manage Manager, err error) {
	// time.Sleep(time.Duration(100) * time.Millisecond)
	if err = sqlDB.First(&manage, "user_name = ?", username).Error; err != nil {
		return
	}
	fmt.Println(manager)
	if err = sqlDB.Model(&manage).Updates(&manager).Error; err != nil {
		return
	}
	return
}

// Reset Password
func (manager *Manager) UpStatusAdmin(status int) {
	sqlDB.Model(&manager).Update("new_status", status)
}

// makePage make page
func makePage(p int) int {
	p = p - 1
	if p <= 0 {
		p = 0
	}
	page := p * 100
	return page
}
