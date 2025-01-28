package service

import (
	"github.com/didsqq/ZeroAgency_Test/internal/models"
	"github.com/didsqq/ZeroAgency_Test/internal/repository"
)

type NewsItemService struct {
	repo repository.News
}

func NewNewsService(repo repository.News) *NewsItemService {
	return &NewsItemService{
		repo: repo,
	}
}

func (s *NewsItemService) GetAll() ([]models.News, error) {
	news, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (s *NewsItemService) Update(newsId int, input models.News) error {
	err := s.repo.Update(newsId, input)
	if err != nil {
		return err
	}

	return nil
}
