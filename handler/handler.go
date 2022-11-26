package handler

import (
	"github.com/camille-plays/coffee2go/dao"
	"github.com/google/uuid"
)

type Handler struct {
	DB dao.Store
}

type TransactionRequest struct {
	Owner      string   `json:"owner"`
	Recipients []string `json:"recipients"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewTestHandler() Handler {
	return Handler{
		DB: &dao.MockStore{
			Users: []*dao.User{
				{ID: uuid.New().String(), Name: "Suzie", Email: "***REMOVED***", Credit: 0},
				{ID: uuid.New().String(), Name: "Camille", Email: "***REMOVED***", Credit: 0},
				{ID: uuid.New().String(), Name: "Mateusz", Email: "***REMOVED***", Credit: 0},
				{ID: uuid.New().String(), Name: "Steve", Email: "***REMOVED***", Credit: 0},
				{ID: uuid.New().String(), Name: "Kenny", Email: "***REMOVED***", Credit: 0},
				{ID: uuid.New().String(), Name: "Brian", Email: "***REMOVED***", Credit: 0},
			},
		},
	}
}
