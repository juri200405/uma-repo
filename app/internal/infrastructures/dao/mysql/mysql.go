package mysql

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type Db *gorm.DB

var dbInstance Db

func GetConnection() Db {
	if dbInstance == nil {
		dbInstance = ConnectDb()
		MigrateDb(dbInstance)
	}
	return dbInstance
}

func ConnectDb() *gorm.DB {
	dataSource := os.Getenv("MYSQL_DATASOURCE")
	maxRetries := 20
	for i := 0; i < maxRetries; i++ {
		db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
		if err != nil {
			log.Printf("failed to connect to database. retrying... (%d/%d) %v", i+1, maxRetries, err)
			time.Sleep(1 * time.Second)
			continue
		}
		return db
	}
	log.Panicln("exceeded max retries of db connecting.")
	return nil
}

func MigrateDb(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Uma{},
		&models.Factor{},
		&models.Race{},
		&models.RaceResult{},
	)
}
