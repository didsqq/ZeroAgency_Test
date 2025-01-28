package repository

import (
	"github.com/didsqq/ZeroAgency_Test/internal/models"
	"gopkg.in/reform.v1"
)

type dbNews struct {
	Id         int64   `reform:"id"`
	Title      string  `reform:"title"`
	Content    string  `reform:"content"`
	Categories []int64 `reform:"categories"`
}

type NewsPostgres struct {
	db *reform.DB
}

func NewNewsPostgres(db *reform.DB) *NewsPostgres {
	return &NewsPostgres{db: db}
}

func (r *NewsPostgres) GetAll() (models.News, error) {
	// var news []models.News
	// getNewsQuery := fmt.Sprintf("SELECT * FROM %s", todoListsTable)
	// err := r.db.FindAllFrom(&news, getNewsQuery)
	// return news, nil
}

func (r *NewsPostgres) Update(newsId int, input models.News) error {

}
