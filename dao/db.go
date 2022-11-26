package dao

type Store interface {
	GetUsers() []*User
	GetUser(id string) *User
	CreateUser(u User)

	GetTransactions() []*Transaction
	GetTransaction(id string) *Transaction
	CreateTransaction(Transaction)
}
