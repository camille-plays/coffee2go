package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storer interface {
	GetUsers() []User
	GetUser(id string) *User
	CreateUser(u *User)
	DeleteUser(u *User) error

	GetTransactions() []Transaction
	GetTransaction(id string) *Transaction
	CreateTransaction(t *Transaction)
}

// Creates a new Store struct for local testing with SQLite
func InitPostgresStore() *Store {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
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
