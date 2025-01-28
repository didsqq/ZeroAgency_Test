package service

import (
	"github.com/didsqq/ZeroAgency_Test/internal/models"
	"github.com/didsqq/ZeroAgency_Test/internal/repository"
)

type News interface {
	GetAll() (models.News, error)
	Update(newsId int, input models.News) error
}

type Service struct {
	News
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		News: NewNewsService(repos.News),
	}
}
