package repository

const (
	queryGetUser = `SELECT id, username, email FROM users WHERE username = $1;`
)
