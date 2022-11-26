package dao

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Credit int    `json:"credit"`
}

type Transaction struct {
	ID         string   `json:"id"`
	Owner      string   `json:"owner"`
	Recipients []string `json:"recipients"`
	Timestamp  int      `json:"timestamp"`
}
