package repository

import (
	"github.com/aszanky/gofolderingproject/internal/repository/payment"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Payment payment.Repository
}

type NewRepositoryParam struct {
	DB *sqlx.DB
}

func NewRepository(param NewRepositoryParam) Repository {
	return Repository{
		Payment: payment.NewPaymentRepository(param.DB),
	}
}
