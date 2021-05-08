package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var TestReturn []interface{}
var TestError []error
var TestIndex int = -1

type Database struct {
	logger *log.Logger
	*sqlx.DB
}

func New(logger *log.Logger) (*Database, error, error) {
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
		TestIndex += 1
		return TestReturn[TestIndex], TestError[TestIndex]
	}
	var temp interface{}
	err := db.Get(&temp, query, parameters)
	return temp, err
}

func (db *Database) Run(query string, parameters ...string) error {
	if os.Getenv("TEST") == "true" {
		db.logger.Printf("returning test result for query: %s with parameters %s\n", query, parameters)
		TestIndex += 1
		return TestError[TestIndex]
	}
	_, err := db.Exec(query, parameters)
	return err
}
