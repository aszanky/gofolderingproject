package usecase

import "github.com/aszanky/gofolderingproject/internal/repository"

type Usecase interface {
}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(
	rep repository.Repository,
) Usecase {
	return &usecase{
		repository: rep,
	}
}
