package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Creates a new Store struct for local testing with SQLite
func InitSqliteStore() *Store {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Transaction{})
	if err != nil {
		panic(err)
	}

	return &Store{
		db: db,
	}
}