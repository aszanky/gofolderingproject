package repository

import "github.com/jmoiron/sqlx"

type Repository interface {
	// CreateUser(ctx context.Context, username string) (err error)
	// GetUser(ctx context.Context, username string) (data model.User, err error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(
	db *sqlx.DB,
) Repository {
	return &repository{
		db: db,
	}
}
