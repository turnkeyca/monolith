package db

import (
	"log"
	"os"
)

var TestReturn []map[string]interface{}
var TestError []error
var TestIndex int = -1

type Database struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Database {
	return &Database{
		logger: logger,
	}
}

func (db *Database) LocalInit() {

}

// func getConnection() interface{} {
// 	return nil
// }

func getSanitizedQueryString(query string, parameters []string) string {
	return ""
}

func (db *Database) Query(query string, parameters ...string) (map[string]interface{}, error) {
	if os.Getenv("TEST") == "true" {
		db.logger.Printf("returning test result for query: %s\n", getSanitizedQueryString(query, parameters))
		TestIndex += 1
		return TestReturn[TestIndex], TestError[TestIndex]
	}
	// conn := getConnection()
	return nil, nil
}

func (db *Database) Put(query string, parameters ...string) error {
	if os.Getenv("TEST") == "true" {
		db.logger.Printf("returning test result for query: %s\n", getSanitizedQueryString(query, parameters))
		TestIndex += 1
		return TestError[TestIndex]
	}
	// conn := getConnection()
	return nil
}

func (db *Database) Delete(query string, parameters ...string) error {
	if os.Getenv("TEST") == "true" {
		db.logger.Printf("returning test result for query: %s\n", getSanitizedQueryString(query, parameters))
		TestIndex += 1
		return TestError[TestIndex]
	}
	// conn := getConnection()
	return nil
}
