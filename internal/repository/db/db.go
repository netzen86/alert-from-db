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

const (
	DBconstring string = "postgres://postgres:alert-from-db@localhost/alert-from-db?sslmode=disable"
)

// DBStorage адрес для подключения к базе данных
type DBStorage struct {
	DB          *sql.DB
	DBconstring string
}

// NewDBStorage функция подключения к базе данных, param = строка для подключения к БД
func NewDBStorage(ctx context.Context, DBconstring string) (*DBStorage, error) {
	var dbstorage DBStorage
	var err error
	dbstorage.DB, err = sql.Open("pgx", DBconstring)
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
		_, err = dbstorage.DB.ExecContext(ctx, stmt, inData.Title, inData.Director)
		log.Println("SQL RES", err)
		if err != nil {
			return fmt.Errorf("insert in table error - %w", err)
		}
	default:
		return fmt.Errorf("insertdata: unknown data type received")
	}

	return nil
}

// функция для вставки данных в базу данных
func (dbstorage *DBStorage) deleteData(ctx context.Context, stmt string, id int64) error {
	var err error
	_, err = dbstorage.DB.ExecContext(ctx, stmt, id)
	log.Println("SQL RES", err)
	if err != nil {
		return fmt.Errorf("insert in table error - %w", err)
	}

	return nil
}

func (dbstorage *DBStorage) CreateTables(ctx context.Context) error {
	var err error

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	stmtFilms := `CREATE TABLE IF NOT EXISTS films 
	("id" SERIAL PRIMARY KEY, "title" TEXT UNIQUE, "director" TEXT); 
	CREATE TABLE IF NOT EXISTS dep1 
	("id" SERIAL PRIMARY KEY, "quarter" TEXT UNIQUE, 
	"plan" integer, "planperc" decimal(5,4) constraint chk_dep1_planperc check (planperc between 0 and 100), 
	"fackt" integer, "facktperc" decimal(5,4) constraint chk_dep1_facktperc check (facktperc between 0 and 100));
	INSERT INTO dep1 (quarter)
	VALUES ('1 квартал'),('2 квартал'),('3 квартал'),('4 квартал');`
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
	ON CONFLICT (title) DO UPDATE 
	  SET director = $2`

	err := dbstorage.insertData(ctx, stmtFilms, f)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("addfilm %w", err)
	}
	return nil
}

func (dbstorage *DBStorage) DelFilmRepo(ctx context.Context, id int64) error {
	stmtFilms := `
	DELETE FROM films
	WHERE id = $1;`

	err := dbstorage.deleteData(ctx, stmtFilms, id)
	if err != nil {
		return fmt.Errorf("delfilmrepo %w", err)
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
