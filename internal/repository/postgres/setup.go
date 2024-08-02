package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func SetUpDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
