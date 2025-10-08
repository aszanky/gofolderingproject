package usecase

import (
	"context"

	"github.com/aszanky/gofolderingproject/internal/model"
	"github.com/aszanky/gofolderingproject/internal/repository"
)

type Usecase interface {
	GetUsers(ctx context.Context, id string) (user model.User, err error)
	IntegrateWithMandiri(ctx context.Context) (err error)
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(
	repo repository.Repository,
) Usecase {
	return &usecase{
		repo: repo,
	}
}
