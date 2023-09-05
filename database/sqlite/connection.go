package sqlite

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteDB struct {
	path                    string
	migrationsDirectoryPath string
}

var (
	LocalSqlite sqliteDB = sqliteDB{
		path:                    "static/db.sqlite",
		migrationsDirectoryPath: "database/sqlite/migration",
	}
)

func (s sqliteDB) Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", s.path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s sqliteDB) Migrate(ctx context.Context) error {
	db, err := s.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	err = filepath.Walk(s.migrationsDirectoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".sql" {
			migration, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			_, err = db.ExecContext(ctx, string(migration))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
