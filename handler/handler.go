package handler

import (
	"github.com/camille-plays/coffee2go/dao"
)

type Handler struct {
	DB dao.Storer
}

type TransactionRequest struct {
	Owner      string   `json:"owner"`
	Recipients []string `json:"recipients"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// Dummy handler used for unit tests
func NewTestHandler() Handler {
	return Handler{
		DB: &dao.MockStore{
			Users: []*dao.User{
				{ID: "96840db2-3676-4399-847e-82e9d2667457", Name: "Suzie", Email: "suzie@transferwise.com", Credit: 0},
				{ID: "b1af6aba-9ec1-4f7b-ad78-d8e4496d9cbe", Name: "Camille", Email: "camille@transferwise.com", Credit: 0},
				{ID: "3b930a43-ab66-48ef-89cd-417dba5d9c8f", Name: "Mateusz", Email: "mateusz@transferwise.com", Credit: 0},
			},
		},
	}
}

func NewLocalHandler() Handler {
	return Handler{
		DB: dao.NewLocalStore(),
	}
}
