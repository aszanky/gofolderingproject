package payment

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aszanky/gofolderingproject/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
)

type paymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(
	db *sqlx.DB,
) Repository {
	return &paymentRepository{
		db: db,
	}
}

func (p *paymentRepository) UpdatePayment(ctx context.Context, username string) (data model.Payment, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Repository_UpdatePayment")
	defer span.Finish()

	//Check if user is already exist
	_, err = p.db.ExecContext(ctx, queryUpdatePayment, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, errors.New("update failed")
		}
		return
	}
	return
}
