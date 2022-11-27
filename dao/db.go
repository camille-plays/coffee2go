package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storer interface {
	GetUsers() []User
	GetUser(id string) *User
	CreateUser(u *User)

	GetTransactions() []Transaction
	GetTransaction(id string) *Transaction
	CreateTransaction(t *Transaction)
}

// Creates a new Store struct for local testing with SQLite
func NewLocalStore() *Store {
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
