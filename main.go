package main

import (
	"fmt"

	"github.com/riyafom/phone-book/contact"
)

func main() {
	c, err := contact.New("Ivan", "Lock", "+79764937655", "email@gmail.com")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	c.SetCompany("Google")

	err = c.AddCategory("Работа")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
	err = c.AddCategory("Home")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
	err = c.RemoveCategory("Рщбота")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
	c.DisplayInfo()

}
