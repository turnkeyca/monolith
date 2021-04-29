package main

import "testing"

func TestGetString(t *testing.T) {
	expected := "Hello world!"
	actual := GetString()
	if actual != expected {
		t.Errorf("actual: '%s', expected: '%s'", actual, expected)
	}
}
