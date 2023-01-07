package dao

import (
	"fmt"

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
	if result := s.db.First(&user); result.Error != nil {
		return nil
	}
	return user
}

func (s *Store) CreateUser(u *User) {
	s.db.Create(u)
}

func (s *Store) DeleteUser(u *User) error {
	if result := s.db.Delete(u); result.RowsAffected < 1 {
		fmt.Println(result.RowsAffected)
		return fmt.Errorf("no user found")
	}

	return nil
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
	s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(t).Error; err != nil {
			return err
		}

		var user = &User{ID: t.Owner}
		if err := tx.First(user).Error; err != nil {
			return err
		}

		if err := tx.Model(user).Update("credit", user.Credit+len(t.Recipients)).Error; err != nil {
			return err
		}

		for _, r := range t.Recipients {
			user.ID = r
			if err := tx.First(user).Error; err != nil {
				return err
			}
			if err := tx.Model(user).Update("credit", user.Credit-1).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
