// database/database.go

package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/thanhtuan260593/todo/models"

	// database driver
	_ "github.com/lib/pq"
)

//Database ..
type Database struct {
	*gorm.DB
	url string
}

//New ...
func New(url string) *Database {
	db := Database{url: url}
	db.initialize()
	db.migrate()
	return &db
}

func (db *Database) initialize() {
	if postgresDB, err := gorm.Open("postgres", db.url); err != nil {
		log.Fatal(err)
	} else {
		db.DB = postgresDB
	}
}

func (db *Database) migrate() {
	db.AutoMigrate(&models.Item{})
}
