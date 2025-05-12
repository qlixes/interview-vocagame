package provider

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

type DatabaseIntf interface {
}

var Db = newDatabase()

var once *sync.Once

func newDatabase() *Database {

	var db *gorm.DB

	once.Do(func() {

		dsn := fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
			Cfg.Db.User,
			Cfg.Db.Pass,
			Cfg.Db.Name,
			Cfg.Db.Host,
			Cfg.Db.Port,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("Unable connect existing database : %s \n", err.Error())
		}
	})

	return &Database{
		Db: db,
	}
}
