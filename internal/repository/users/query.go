package users

const (
	queryGetUsers = "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
)
