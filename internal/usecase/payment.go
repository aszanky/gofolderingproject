package usecase

import (
	"github.com/aszanky/gofolderingproject/internal/repository/payment"
)

type paymentUsecase struct {
	paymentRepository payment.Repository
}

func NewPaymentUsecase(
	paymentRepo payment.Repository,
) PaymentDomain {
	return &paymentUsecase{
		paymentRepository: paymentRepo,
	}
}

func (p *paymentUsecase) IntegrateWithMandiri() (err error) {
	return err
}
