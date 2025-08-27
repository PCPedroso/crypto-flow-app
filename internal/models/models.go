package models

type Transaction struct {
    ID          string  `json:"id"`
    UserID      string  `json:"user_id"`
    Amount      float64 `json:"amount"`
    Currency    string  `json:"currency"`
    Timestamp   int64   `json:"timestamp"`
}

type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Balance  float64 `json:"balance"`
}