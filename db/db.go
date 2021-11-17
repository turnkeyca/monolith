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
	db, errOpen := sqlx.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONN"))
	if errOpen != nil {
		return nil, errOpen
	}
	errPing := db.Ping()
	if errPing != nil {
		return nil, errPing
	}
	logger.Println(os.Getenv("DB_CONN"))
	return &Database{
		logger: logger,
		DB:     db,
	}, nil
}

func (db *Database) Run(query string, parameters ...interface{}) error {
	_, err := db.Exec(query, parameters...)
	return err
}
