package db

import (
	"fmt"
	"os"
	"regexp"
)

var TestReturn []interface{}
var TestError []error
var TestQuery []string

func (db *Database) SetNextTestReturn(next interface{}) {
	if os.Getenv("TEST") != "true" {
		panic("not implemented for non tests!")
	}
	if len(TestReturn) == 0 {
		TestReturn = []interface{}{}
	}
	TestReturn = append(TestReturn, next)
}

func (db *Database) getNextTestReturn() interface{} {
	if os.Getenv("TEST") != "true" {
		panic("not implemented for non tests!")
	}
	if len(TestReturn) == 0 {
		return nil
	}
	next := TestReturn[0]
	TestReturn = TestReturn[1:]
	return next
}

func (db *Database) SetNextTestError(next error) {
	if os.Getenv("TEST") != "true" {
		panic("not implemented for non tests!")
	}
	if len(TestError) == 0 {
		TestError = []error{}
	}
	TestError = append(TestError, next)
}

func (db *Database) getNextTestError() error {
	if os.Getenv("TEST") != "true" {
		panic("not implemented for non tests!")
	}
	if len(TestError) == 0 {
		return nil
	}
	next := TestError[0]
	TestError = TestError[1:]
	return next
}

func (db *Database) GetNextTestQuery() string {
	if os.Getenv("TEST") != "true" {
		panic("not implemented for non tests!")
	}
	if len(TestQuery) == 0 {
		return ""
	}
	next := TestQuery[0]
	TestQuery = TestQuery[1:]
	return next
}

func pushQuery(query string, parameters ...string) {
	if os.Getenv("TEST") != "true" {
		panic("not implemented for non tests!")
	}
	if len(TestQuery) == 0 {
		TestQuery = []string{}
	}
	next := query
	for i, param := range parameters {
		re := regexp.MustCompile(fmt.Sprintf(`\$%d`, i+1))
		next = re.ReplaceAllString(next, param)
	}
	TestQuery = append(TestQuery, next)
}
