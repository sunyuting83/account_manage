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
	sqlDB.AutoMigrate(&Users{}, &Projects{}, &Accounts{}, &Filed{}, &Manager{})

	var (
		manager *Manager
	)

	if m := sqlDB.First(&manager); m.Error != nil {
		if m.Error.Error() == "record not found" {
			u := Manager{
				UserName:  "admin",
				Password:  pwd,
				NewStatus: 0,
			}
			sqlDB.Create(&u)
		}
	}
}
