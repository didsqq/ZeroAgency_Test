package repository

import (
	"fmt"

	"github.com/didsqq/ZeroAgency_Test/internal/models"
	"gopkg.in/reform.v1"
)

type NewsPostgres struct {
	db *reform.DB
}

func NewNewsPostgres(db *reform.DB) *NewsPostgres {
	return &NewsPostgres{db: db}
}

func (r *NewsPostgres) GetAll() ([]models.News, error) {
	var newsSl []models.News
	query := fmt.Sprintf("SELECT * FROM %s", newsTable)
	rows, err := r.db.Query(query, nil)

	for rows.Next() {
		var news models.News
		err := rows.Scan(&news.Id, &news.Title, &news.Content, &news.Categories)
		if err != nil {
			return nil, err
		}
		newsSl = append(newsSl, news)
	}

	if err != nil {
		return nil, err
	}

	return newsSl, nil
}

func (r *NewsPostgres) Update(newsId int, input models.News) error {
	query := fmt.Sprintf("UPDATE %s SET title = $1, content = $2, categories = NOW() WHERE id = $3", newsTable)

	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
