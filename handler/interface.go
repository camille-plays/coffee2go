package handler

type TransactionRequest struct {
	Owner      string   `json:"owner"`
	Recipients []string `json:"recipients"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
