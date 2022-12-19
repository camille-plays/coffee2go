package dao

import (
	"gorm.io/gorm"
  	"gorm.io/driver/postgres"
)

type Storer interface {
	GetUsers() []User
	GetUser(id string) *User
	CreateUser(u *User)
	DeleteUser(u *User)

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