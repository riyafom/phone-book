package contact

import (
	"testing"
)

func TestFullName(t *testing.T) {
	contact := Contact{
		FirstName: "Alex",
		LastName:  "Gremory",
	}

	res := contact.FullName()
	expected := "Alex Gremory"

	if res != expected {
		t.Errorf("Результат: %s не соответствует ожидаемому: %s", res, expected)
	}
}
