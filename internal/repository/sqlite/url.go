package sqlite

import (
	dbSQL "database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/iwerqfx/url-shortener/internal/model"
)

const urlsTableName = "urls"

type URLRepository interface {
	Create(url, alias string) error
	GetByAlias(alias string) (model.URL, error)
	IncreaseViews(alias string) error
}
type urlRepository struct {
	*Repository
}

func NewURLRepository(repository *Repository) URLRepository {
	stmt, err := repository.db.Prepare(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s(
		    id INTEGER PRIMARY KEY,
		    url TEXT NOT NULL,
		    alias TEXT NOT NULL UNIQUE,
		    views INTEGER DEFAULT 0
		);
		CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias);
	`, urlsTableName))
	if err != nil {
		panic("error preparing statement for creating urls table: " + err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		panic("error executing statement for creating urls table: " + err.Error())
	}

	return &urlRepository{
		repository,
	}
}

func (r *urlRepository) Create(url, alias string) error {
	sql, args, err := r.sb.Insert("urls").
		Columns("url", "alias").
		Values(url, alias).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sql, args...)

	return err
}

func (r *urlRepository) GetByAlias(alias string) (model.URL, error) {
	sql, args, err := r.sb.Select("*").
		From("urls").
		Where(squirrel.Eq{"alias": alias}).
		ToSql()
	if err != nil {
		return model.URL{}, err
	}

	var url model.URL
	if err = r.db.QueryRow(sql, args...).Scan(&url.ID, &url.URL, &url.Alias, &url.Views); err != nil {
		if errors.Is(err, dbSQL.ErrNoRows) {
			return model.URL{}, model.ErrURLNotFound
		}

		return model.URL{}, err
	}

	return url, nil
}

func (r *urlRepository) IncreaseViews(alias string) error {
	sql, args, err := r.sb.Update(urlsTableName).
		Set("views", squirrel.Expr("views + ?", 1)).
		Where(squirrel.Eq{"alias": alias}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sql, args...)

	return err
}
