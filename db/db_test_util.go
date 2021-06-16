package db

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var TestReturn [][]interface{}
var TestError []error
var TestQuery []string

func (db *Database) SetNextTestReturn(next []interface{}) {
	if os.Getenv("TEST") != "true" {
		panic("not implemented for non tests!")
	}
	if len(TestReturn) == 0 {
		TestReturn = [][]interface{}{}
	}
	TestReturn = append(TestReturn, next)
}

func (db *Database) GetNextTestReturn() []interface{} {
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

func (db *Database) GetNextTestError() error {
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

func (db *Database) PushQuery(query string, parameters ...interface{}) {
	if os.Getenv("TEST") != "true" {
		panic("not implemented for non tests!")
	}
	if len(TestQuery) == 0 {
		TestQuery = []string{}
	}
	next := query
	for i, param := range parameters {
		re := regexp.MustCompile(fmt.Sprintf(`\$%d`, i+1))
		next = re.ReplaceAllString(next, getString(param))
	}
	TestQuery = append(TestQuery, next)
}

func getString(obj interface{}) string {
	switch v := obj.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', 2, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return obj.(string)
	}
}
