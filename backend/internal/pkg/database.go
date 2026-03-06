package pkg

import (
	"fmt"
	"log"

	"petService/backend/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	var db *gorm.DB
	var err error

	if cfg.DBType == "postgres" || cfg.DBType == "supabase" {
		// PostgreSQL 连接方式（支持直接传入标准完整连接串）
		// 如果在本地配了完整的 DB_HOST (如 postgres://user:pwd@host:port/db)
		// 我们直接把它拿来做 DSN，如果没有包含前缀，说明传的只有地址本身
		dsn := cfg.DBHost
		if len(dsn) < 10 || dsn[0:10] != "postgresql" {
			dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai",
				cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
		}
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("PostgreSQL 数据库连接失败: %v", err)
		}
	} else {
		// 默认 MySQL 连接方式 (保留原先的设定)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("MySQL 数据库连接失败: %v", err)
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取底层数据库连接失败: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("数据库连接成功")
	return db
}
