package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	logger *log.Logger
	*sqlx.DB
}

func New(logger *log.Logger) (*Database, error) {
	if os.Getenv("TEST") == "true" {
		return &Database{
			logger: logger,
		}, nil
	}
	db, errOpen := sqlx.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONN"))
	if errOpen != nil {
		return nil, errOpen
	}
	errPing := db.Ping()
	if errPing != nil {
		return nil, errPing
	}
	return &Database{
		logger: logger,
		DB:     db,
	}, nil
}

func (db *Database) Run(query string, parameters ...interface{}) error {
	if os.Getenv("TEST") == "true" {
		db.logger.Printf("returning test result for query: %s with parameters %s\n", query, parameters)
		db.PushQuery(query, parameters...)
		return db.GetNextTestError()
	}
	_, err := db.Exec(query, parameters...)
	return err
}
