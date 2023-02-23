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

func UserCheckUserName(username string) (user *Users, err error) {
	if err = sqlDB.First(&user, "user_name = ? ", username).Error; err != nil {
		return
	}
	return
}

func (user *Users) Insert() (err error) {
	sqlDB.Create(&user)
	return nil
}

// Delete Admin
func (user *Users) DeleteOne(id int64) {
	// time.Sleep(time.Duration(100) * time.Millisecond)
	sqlDB.Where("id = ?", id).Delete(&user)
}

// Check ID
func UserCheckID(id int64) (user *Users, err error) {
	if err = sqlDB.First(&user, "id = ?", id).Error; err != nil {
		return
	}
	return
}

// Get Count
func (user *Users) GetCount() (count int64, err error) {
	if err = sqlDB.Model(&user).Count(&count).Error; err != nil {
		return
	}
	return
}

// Admin List
func GetUsersList(page int) (user *Users, err error) {
	p := makePage(page)
	if err = sqlDB.
		Select("id, user_name, new_status, created_at").
		Order("id desc").
		Limit(100).Offset(p).
		Find(&user).Error; err != nil {
		return
	}
	return
}

// Reset Password
func (user *Users) UserResetPassword(username string) (users Users, err error) {
	// time.Sleep(time.Duration(100) * time.Millisecond)
	if err = sqlDB.First(&users, "user_name = ?", username).Error; err != nil {
		return
	}
	if err = sqlDB.Model(&user).Updates(&users).Error; err != nil {
		return
	}
	return
}

// Reset Password
func (user *Users) UserUpStatusAdmin(status int) {
	sqlDB.Model(&user).Update("new_status", status)
}
