package backend

import (
	"alert-from-db/internal/repository/db"
	"context"
	"fmt"
)

// MakeDBMigrations создание таблицы films
func MakeDBMigrations(ctx context.Context, storage *db.DBStorage) error {
	err := storage.CreateTables(ctx)
	if err != nil {
		return fmt.Errorf("error when creating tables %v", err)
	}
	return nil
}
