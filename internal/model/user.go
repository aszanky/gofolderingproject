package model

type AuthPayload struct {
	Username string `json:"username" binding:"required"`
}

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type AllUser struct {
	Users []User
}
