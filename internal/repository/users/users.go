package users

import (
	"context"

	"github.com/aszanky/gofolderingproject/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
)

type Repository interface {
	GetUsers(ctx context.Context, id string) (user model.User, err error)
}

type usersRepository struct {
	db *pgxpool.Pool
}

func NewUsersRepository(
	db *pgxpool.Pool,
) Repository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) GetUsers(ctx context.Context, id string) (user model.User, err error) {
	tracer := otel.Tracer("repository-layer")
	ctx, span := tracer.Start(ctx, "repository.GetUsers")
	defer span.End()

	err = r.db.QueryRow(ctx, queryGetUsers, id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}
