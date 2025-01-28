package repository

import (
	"github.com/didsqq/ZeroAgency_Test/internal/models"
	"gopkg.in/reform.v1"
)

type News interface {
	GetAll() ([]models.News, error)
	Update(newsId int, input models.News) error
}

type Repository struct {
	News
}

func NewRepository(db *reform.DB) *Repository {
	return &Repository{
		News: NewNewsPostgres(db),
	}
}
