package usecase

import (
	"github.com/aszanky/gofolderingproject/internal/repository"
	"github.com/jmoiron/sqlx"
)

type Usecase struct {
	Payment PaymentDomain
}

type NewUsecaseParam struct {
	DB *sqlx.DB
}

func NewService(param NewUsecaseParam) Usecase {
	repo := repository.NewRepository(repository.NewRepositoryParam{
		DB: param.DB,
	})
	return Usecase{
		Payment: NewPaymentUsecase(repo.Payment),
	}
}

type PaymentDomain interface {
	IntegrateWithMandiri() (err error)
}
