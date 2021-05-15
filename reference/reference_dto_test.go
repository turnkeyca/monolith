package reference

import (
	"testing"
)

func TestValidate(t *testing.T) {
	dto := New()
	err := dto.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
