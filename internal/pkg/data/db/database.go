package db

import (
	"fmt"
	"time"
	"vanir/internal/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func SetupDB() {
	var db = DB

	conf := config.GetConfig()

	driver := conf.Database.Driver
	database := conf.Database.DbName

	username := conf.Database.Username
	password := conf.Database.Password
	host := conf.Database.Host
	port := conf.Database.Port

	if driver == "sqlite" { // SQLITE
		db, err = gorm.Open(sqlite.Open(database), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			fmt.Println("error connecting to db: ", err)
		}
	} else if driver == "postgres" {
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=America/Sao_Paulo", host, username, password, database, port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			fmt.Println("error connecting to db: ", err)
		}
	} else { // Not implemented
		err = fmt.Errorf("other drivers not implemented")
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("error setting up db: ", err)
	}
	sqlDB.SetMaxIdleConns(conf.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.Database.MaxLifetime) * time.Second)
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
