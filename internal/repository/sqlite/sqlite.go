package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
)

type Repository struct {
	log *slog.Logger
	db  *sql.DB
	sb  squirrel.StatementBuilderType
}

func NewRepository(log *slog.Logger, db *sql.DB) *Repository {
	sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)
	return &Repository{
		log: log,
		db:  db,
		sb:  sb,
	}
}

func NewDB(url string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		return nil, fmt.Errorf("error opening sqlite3 db: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging sqlite3 db: %w", err)
	}

	return db, nil
}
