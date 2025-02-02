// Package db - пакет для работы с хранилищем типа база данных
package db

import (
	"alert-from-db/internal/data/films"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// DBStorage адрес для подключения к базе данных -  "postgres://postgres:collectmetrics@localhost/collectmetrics?sslmode=disable"
type DBStorage struct {
	DB          *sql.DB
	DBconstring string
}

// NewDBStorage функция подключения к базе данных, param = строка для подключения к БД
func NewDBStorage(ctx context.Context, param string) (*DBStorage, error) {
	var dbstorage DBStorage
	var err error
	dbstorage.DB, err = sql.Open("pgx", param)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to data base %w", err)
	}
	err = dbstorage.DB.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot ping data base %w", err)
	}
	return &dbstorage, nil
}

// функция для вставки данных в базу данных
func (dbstorage *DBStorage) insertData(ctx context.Context, stmt string, inData interface{}) error {
	var err error
	switch inData := inData.(type) {
	case films.Film:
		_, err = dbstorage.DB.ExecContext(ctx, stmt, inData)
		if err != nil {
			return fmt.Errorf("insert in table error - %w", err)
		}
	default:
		return fmt.Errorf("insertdata: unknown data type received")
	}

	return nil
}

func (dbstorage *DBStorage) CreateTables(ctx context.Context) error {
	var err error

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	stmtFilms := `CREATE TABLE IF NOT EXISTS gauge 
	("id" SERIAL PRIMARY KEY, "title" TEXT UNIQUE, "director" TEXT)`
	_, err = dbstorage.DB.ExecContext(ctx, stmtFilms)
	if err != nil {
		return fmt.Errorf("create table error - %w", err)
	}
	return nil
}

func (dbstorage *DBStorage) AddFilmRepo(ctx context.Context, f films.Film) error {
	stmtFilms := `
	INSERT INTO films (title, director) 
	VALUES ($1, $2)
	ON CONFLICT (name) DO UPDATE 
	  SET value = $2`

	err := dbstorage.insertData(ctx, stmtFilms, f)
	if err != nil {
		return fmt.Errorf("gauge %w", err)
	}
	return nil
}

func (dbstorage *DBStorage) DelFilmRepo(ctx context.Context, id int64) error {
	stmtFilms := `
	DELETE FROM films
	WHERE id = $1;`

	err := dbstorage.insertData(ctx, stmtFilms, id)
	if err != nil {
		return fmt.Errorf("gauge %w", err)
	}
	return nil
}

func (dbstorage *DBStorage) GetFilmsRepo(ctx context.Context) ([]films.Film, error) {
	var f []films.Film
	var err error

	smtp := `
	SELECT id, Title, director
	FROM films;`

	rows, err := dbstorage.DB.QueryContext(ctx, smtp)
	if err != nil {
		return nil, fmt.Errorf("getfilmsrepo: error when execute select %w", err)
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Printf("getfilmsrepo: error when close rows %v", err)
		}
	}()
	for rows.Next() {
		var id int64
		var title string
		var director string

		err = rows.Scan(&id, &title, &director)
		if err != nil {
			return nil, fmt.Errorf("getfilmsrepo: error scan %w", err)
		}
		f = append(f, films.Film{ID: id, Title: title, Director: director})
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("errors rows %w", err)
	}
	return f, nil
}
