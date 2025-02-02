// Package repositories - пакет содержит интерфейс для работы с хранилищами
package repositories

import (
	"alert-from-db/internal/data/films"
	"context"
)

type Repo interface {
	AddFilmRepo(ctx context.Context, film films.Film) error
	DelFilmRepo(ctx context.Context, FilmID int64) error
	GetFilmsRepo(ctx context.Context) ([]films.Film, error)
}
