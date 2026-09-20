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

func TestIsSameContact(t *testing.T) {
	contact1 := Contact{
		FirstName: "Александра",
		LastName:  "Гремори",
		phone:     "+79991112233",
	}
	contact2 := Contact{
		FirstName: "Анна",
		LastName:  "Иванова",
		phone:     "+79991112233",
	}

	result := contact1.IsSameContact(contact2)
	expected := true
	if result != expected {
		t.Errorf("Результат: %v не соответствует ожидаемому: %v", result, expected)
	}
}

func TestSetPhoneInvalid(t *testing.T) {
	var contact Contact
	phones := []string{"adc", "123", "1+23456789"}

	for i, phone := range phones {
		result := contact.SetPhone(phone)

		if result == nil {
			t.Errorf("тест %d: для телефона %q ожидалась ошибка", i+1, phone)
		}

	}
}

func TestSetPhoneValid(t *testing.T) {
	var contact Contact
	phones := []string{"+123456789", "8123456789", "7655565545"}

	for i, phone := range phones {
		result := contact.SetPhone(phone)

		if result != nil {
			t.Errorf("тест %d: для телефона %q ожидалось nil", i+1, phone)
		}
	}
}

func TestSetPhone(t *testing.T) {
	var contact Contact
	phone := "+123456789"

	contact.SetPhone(phone)

	if contact.phone != phone {
		t.Errorf("contact.phone = %q ожидалось получить %q", contact.phone, phone)
	}
}
