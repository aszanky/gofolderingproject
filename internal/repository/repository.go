package repository

import (
	"github.com/aszanky/gofolderingproject/internal/repository/payment"
	"github.com/aszanky/gofolderingproject/internal/repository/users"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Payment payment.Repository
	Users   users.Repository
}

type NewRepositoryParam struct {
	DB    *sqlx.DB
	PGXDB *pgxpool.Pool
}

func NewRepository(param NewRepositoryParam) Repository {
	return Repository{
		Payment: payment.NewPaymentRepository(param.DB),
		Users:   users.NewUsersRepository(param.PGXDB),
	}
}
