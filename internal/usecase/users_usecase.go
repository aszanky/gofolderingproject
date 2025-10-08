package usecase

import (
	"context"

	"github.com/aszanky/gofolderingproject/internal/model"
	"go.opentelemetry.io/otel"
)

func (u *usecase) GetUsers(ctx context.Context, id string) (user model.User, err error) {
	tracer := otel.Tracer("usecase-layer")
	ctx, span := tracer.Start(ctx, "usecase.GetUsers")
	defer span.End()

	user, err = u.repo.Users.GetUsers(ctx, id)
	if err != nil {
		return user, err
	}
	return user, nil
}
