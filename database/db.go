package database

import (
	"AccountManage/utils"
	"database/sql"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Eloquent *sql.DB
	sqlDB    *gorm.DB
)

// InitDB init db
func InitDB(pwd string, confYaml *utils.Config) {
	DNString := strings.Join([]string{"host=", confYaml.DBHost, " user=", confYaml.Username, " password=", confYaml.Password, " dbname=", confYaml.DBName, " port=", confYaml.DBProt, " sslmode=disable TimeZone=Asia/Shanghai"}, "")
	sqlDB, _ = gorm.Open(postgres.New(postgres.Config{
		DSN:                  DNString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	Eloquent, _ = sqlDB.DB()
	Eloquent.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	Eloquent.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	Eloquent.SetConnMaxLifetime(time.Hour)
	sqlDB.AutoMigrate(&Users{}, &Projects{}, &Accounts{}, &Filed{}, &ManageGroups{}, &Manager{}, &Power{})

	var (
		group   *ManageGroups
		manager *Manager
		GroupID uint
		power   = []Power{
			{TopID: 0,
				Title:     "系统菜单管理",
				NewStatus: 0,
				Api:       "systemenu"},
			{TopID: 1,
				Title:     "添加顶级菜单",
				NewStatus: 1,
				Api:       "addtopmenu"},
			{TopID: 1,
				Title:     "添加下级菜单",
				NewStatus: 2,
				Api:       "addsmenu"},
			{TopID: 1,
				Title:     "删除下级菜单",
				NewStatus: 2,
				Api:       "systemenu"},
			{TopID: 1,
				Title:     "添加功能",
				NewStatus: 3,
				Api:       "addfunction"},
			{TopID: 1,
				Title:     "删除菜单",
				NewStatus: 3,
				Api:       "delsystem",
			},
		}
	)
	if g := sqlDB.First(&group); g.Error != nil {
		if g.Error.Error() == "record not found" {
			u := ManageGroups{
				Title: "超级管理组",
				Power: "",
			}
			sqlDB.Create(&u)
			GroupID = u.ID
		}
	}

	if m := sqlDB.First(&manager); m.Error != nil {
		if m.Error.Error() == "record not found" {
			u := Manager{
				ManageGroupsID: GroupID,
				UserName:       "admin",
				Password:       pwd,
				NewStatus:      0,
			}
			sqlDB.Create(&u)
		}
	}
}
