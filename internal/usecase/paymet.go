package usecase

import (
	"github.com/aszanky/gofolderingproject/internal/repository/payment"
)

type paymentService struct {
	paymentRepository payment.Repository
}

func NewPaymentService(
	paymentRepo payment.Repository,
) PaymentDomain {
	return &paymentService{
		paymentRepository: paymentRepo,
	}
}

func (p *paymentService) IntegrateWithMandiri() (err error) {
	return err
}
