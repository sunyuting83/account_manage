package database

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Eloquent *sql.DB
	sqlDB    *gorm.DB
)

// InitDB init db
func InitDB() {
	sqlDB, _ = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=192.168.1.90 user=acc_manage password=123456 dbname=acc_manage port=5432 sslmode=disable TimeZone=Asia/Shanghai",
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
	)
	g := sqlDB.First(&group)
	fmt.Println(g)
	m := sqlDB.First(&manager)
	fmt.Println(m)
}
