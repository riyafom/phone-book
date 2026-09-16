package contact

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"
)

type Contact struct {
	FirstName string
	LastName  string
	Company   string

	phone      string
	email      string
	createdAt  time.Time
	categories []string
}

func New(firstName, lastName, phone, email string) (*Contact, error) {
	if strings.TrimSpace(firstName) == "" {
		return nil, errors.New("имя не может быть пустым")
	}
	if strings.TrimSpace(lastName) == "" {
		return nil, errors.New("фамилия не может быть пустой")
	}

	if !isValidPhone(phone) {
		return nil, errors.New("некорректный номер телефона")
	}

	if !isValidEmail(email) {
		return nil, errors.New("некорректный email")
	}

	return &Contact{
		FirstName:  firstName,
		LastName:   lastName,
		Company:    "",
		phone:      phone,
		email:      email,
		createdAt:  time.Now(),
		categories: []string{},
	}, nil

}

func isValidPhone(phone string) bool {

	cleaned := strings.ReplaceAll(phone, " ", "")

	if len(cleaned) < 10 {
		return false
	}

	pattern := `^\+?\d+$`
	matched, _ := regexp.MatchString(pattern, cleaned)
	return matched
}

func isValidEmail(email string) bool {
	cleaned := strings.TrimSpace(email)

	pattern := `^[^@\s]+@[^@\s]+\.[^@\s]+$`
	matched, _ := regexp.MatchString(pattern, cleaned)

	return matched
}

func (c *Contact) Phone() string {
	return c.phone
}

func (c *Contact) Email() string {
	return c.email
}

func (c *Contact) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Contact) Categories() []string {
	if len(c.categories) == 0 {
		return []string{}
	}

	result := make([]string, len(c.categories))
	copy(result, c.categories)
	return result
}

func (c *Contact) SetPhone(phone string) error {
	if !isValidPhone(phone) {
		return errors.New("некорректный номер телефона")
	}

	c.phone = phone
	return nil
}

func (c *Contact) SetEmail(email string) error {
	if !isValidEmail(email) {
		return errors.New("некорректный email")
	}

	c.email = email
	return nil
}

func (c *Contact) SetCompany(company string) {
	c.Company = company
}

func (c *Contact) AddCategory(category string) error {
	cleaned := strings.TrimSpace(category)
	if cleaned == "" {
		return errors.New("категория не может быть пустой")
	}

	if c.HasCategory(cleaned) {
		return fmt.Errorf("категория '%s' уже существует", cleaned)
	}
	c.categories = append(c.categories, cleaned)
	return nil
}

func (c *Contact) HasCategory(category string) bool {
	if slices.Contains(c.categories, category) {
		return true
	}
	return false
}

func (c *Contact) RemoveCategory(category string) error {
	cleaned := strings.TrimSpace(category)

	if !c.HasCategory(cleaned) {
		return errors.New("категория не найдена")
	}

	index := slices.Index(c.categories, cleaned)
	c.categories = append(c.categories[:index], c.categories[index+1:]...)
	return nil
}

func (c Contact) FullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

func (c *Contact) DisplayInfo() {
	fmt.Println("============================")
	fmt.Printf("Контакт: %s\n", c.FullName())
	fmt.Println("============================")

	fmt.Printf("Компания: %s\n", c.Company)
	fmt.Printf("Телефон: %s\n", c.phone)
	fmt.Printf("Email: %s\n", c.email)
	fmt.Printf("Категории: %s\n", strings.Join(c.categories, ", "))

	fmt.Printf("Создан: %v\n", c.createdAt.Format("2006-01-02 15:04:05"))

	fmt.Println("============================")

}

func (c Contact) IsSameContact(other Contact) bool {
	if c.phone == other.phone {
		return true
	}
	return false
}

func (c *Contact) UpdateContact(phone, email, company string) error {
	if !isValidPhone(phone) {
		return errors.New("некорректный номер телефона")
	}
	if !isValidEmail(email) {
		return errors.New("некорректный email")
	}

	c.phone = phone
	c.email = email
	c.Company = company
	return nil
}
