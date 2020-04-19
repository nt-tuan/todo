// infrastructure/database.go

package database

import (
	"log"

	"github.com/jinzhu/gorm"

	// database driver
	_ "github.com/lib/pq"

	"github.com/thanhtuan260593/todo/core/entities"
)

//Database ..
type Database struct {
	*gorm.DB
	URL string
}

//New ...
func New(url string) *Database {
	db := Database{URL: url}
	db.initialize()
	db.migrate()
	return &db
}

func (db *Database) initialize() {
	if postgresDB, err := gorm.Open("postgres", db.URL); err != nil {
		log.Fatal(err)
	} else {
		db.DB = postgresDB
	}
}

func (db *Database) migrate() {
	db.AutoMigrate(&entities.Item{})
}
