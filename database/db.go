package database

import (
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Eloquent *sql.DB
)

// InitDB init db
func InitDB() {
	sqlDB, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=192.168.1.90 user=acc_manage password=123456 dbname=acc_manage port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	Eloquent, _ = sqlDB.DB()
	Eloquent.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	Eloquent.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	Eloquent.SetConnMaxLifetime(time.Hour)
}
