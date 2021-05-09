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

func New(logger *log.Logger) (*Database, error, error) {
	if os.Getenv("TEST") == "true" {
		return &Database{
			logger: logger,
		}, nil, nil
	}
	db, errOpen := sqlx.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONN"))
	errPing := db.Ping()
	return &Database{
		logger: logger,
		DB:     db,
	}, errOpen, errPing
}

func (db *Database) Query(query string, parameters ...string) (interface{}, error) {
	if os.Getenv("TEST") == "true" {
		db.logger.Printf("returning test result for query: %s with parameters %s\n", query, parameters)
		pushQuery(query, parameters...)
		return db.getNextTestReturn(), db.getNextTestError()
	}
	var temp interface{}
	err := db.Get(&temp, query, parameters)
	return temp, err
}

func (db *Database) Run(query string, parameters ...string) error {
	if os.Getenv("TEST") == "true" {
		db.logger.Printf("returning test result for query: %s with parameters %s\n", query, parameters)
		pushQuery(query, parameters...)
		return db.getNextTestError()
	}
	_, err := db.Exec(query, parameters)
	return err
}
