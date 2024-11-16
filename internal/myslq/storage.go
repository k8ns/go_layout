package myslq

import (
	"article/internal/article"
	"database/sql"
)

type Storage struct {
    db *sql.DB
}


func (s *Storage) Add(a *article.Article) (int, error) {
    return 1, nil
}
