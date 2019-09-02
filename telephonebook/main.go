package main

import (
	"fmt"
)

// Contact - телефонная книга
type Contact struct {
	Surname string
	Name    string
	Company string
	Number  string
}

//TelephoneBook - массив с контактами
var TelephoneBook map[string]Contact

//C - переменная структуры контакта
var C Contact

//Action - переменная значения выбора
var Action string

//Surname - переменная фамилии
var Surname string

//Find -переменная фамилии для поиска в карте
var Find string

// WriteIn - добавление контакта в телефонную книгу
func (c *Contact) WriteIn() {

	c.Surname = Surname
	fmt.Println("Введите Имя:")
	fmt.Scanln(&c.Name)
	fmt.Println("Введите название организации:")
	fmt.Scanln(&c.Company)
	fmt.Println("Введите телефонный номер:")
	fmt.Scanln(&c.Number)

}

//Options - функция задающая параметры добавления в телефонную книгу, а так же поиска в ней
func Options() {
	fmt.Println("Выберите действие:\n 1 - поиск контакта\n 2 - добавление нового\n 3 - Выход")
	fmt.Scanln(&Action)
	if Action == "1" {
		fmt.Println("Поиск контактов\nВведите фамилию:")
		fmt.Scanln(&Find)
		fmt.Println(TelephoneBook[Find])
		Options()
	} else {
		if Action == "2" {
			fmt.Println("Ввдите Фамилию:")
			fmt.Scanln(&Surname)
			C.WriteIn()
			TelephoneBook[Surname] = C
			fmt.Println(TelephoneBook)
			Options()
		} else {
			if Action == "3" {
				fmt.Println("Спасибо за внимание!")
			} else {
				fmt.Println("Неправильно заданы значения!")
				Options()
			}
		}
	}
}

func main() {
	TelephoneBook = make(map[string]Contact)
	Options()

}
