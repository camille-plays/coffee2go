package dao

import "github.com/lib/pq"

type User struct {
	ID     string `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Credit int    `json:"credit"`
}

type Transaction struct {
	ID         string         `json:"id" gorm:"primary_key"`
	Owner      string         `json:"owner"`
	Recipients pq.StringArray `json:"recipients" gorm:"type:text[]"`
	Timestamp  int            `json:"timestamp"`
}
