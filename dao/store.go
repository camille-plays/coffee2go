package dao

import (
	"gorm.io/gorm"
)

// Store implements the Storer interface with gorm
type Store struct {
	db *gorm.DB
}

func (s *Store) GetUsers() []User {
	var users []User
	s.db.Find(&users)
	return users
}

func (s *Store) GetUser(id string) *User {
	var user = &User{ID: id}
	s.db.First(user)
	return user
}

func (s *Store) CreateUser(u *User) {
	s.db.Create(u)
}

func (s *Store) GetTransactions() []Transaction {
	var tr []Transaction
	s.db.Find(&tr)
	return tr
}

func (s *Store) GetTransaction(id string) *Transaction {
	var tr = &Transaction{ID: id}
	s.db.First(tr)
	return tr
}

func (s *Store) CreateTransaction(t *Transaction) {
	s.db.Create(t)
	// TODO: increment credit of user by 1
	s.db.Model(&User{ID: t.Owner}).Update("credit", 1)
}
