package dao

type MockStore struct {
	Users        []*User
	Transactions []Transaction
}

func (s *MockStore) GetUsers() []User {
	users := make([]User, len(s.Users))
	for k, v := range s.Users {
		users[k] = *v
	}
	return users
}

func (s *MockStore) GetUser(id string) *User {
	for _, v := range s.Users {
		if v.ID == id {
			return v
		}
	}
	return nil
}

func (s *MockStore) CreateUser(u *User) {
	s.Users = append(s.Users, u)
}

func (s *MockStore) DeleteUser(u *User) error {
	i := -1
	for k, v := range s.Users {
		if v == u {
			i = k
		}
	}
	s.Users = append(s.Users[:i], s.Users[i+1:]...)

	return nil
}

func (s *MockStore) GetTransactions() []Transaction {
	return s.Transactions
}

func (s *MockStore) GetTransaction(id string) *Transaction {
	for _, v := range s.Transactions {
		if v.ID == id {
			return &v
		}
	}
	return nil
}

func (s *MockStore) CreateTransaction(t *Transaction) {
	// persist transaction history
	s.Transactions = append(s.Transactions, *t)

	// increment credit of owner
	for _, a := range s.Users {
		if t.Owner == a.ID {
			a.Credit += len(t.Recipients)
		}
	}

	// decrement credit of each recipient
	for _, r := range t.Recipients {
		for _, u := range s.Users {
			if r == u.ID {
				u.Credit -= 1
			}
		}
	}
}
