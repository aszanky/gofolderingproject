package payment

import (
	"context"

	"github.com/aszanky/gofolderingproject/internal/model"
)

type Repository interface {
	UpdatePayment(ctx context.Context, username string) (data model.Payment, err error)
}
