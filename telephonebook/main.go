package main

import (
	"fmt"
	"sort"
)

// Contact - телефонная книга
type Contact struct {
	Surname string
	Name    string
	Company string
	Number  string
}

//sortSurname - тип для сортировки по фамилии
type sortSurname []Contact

//TelephoneBook - массив с контактами
var TelephoneBook []Contact

//C - переменная структуры контакта
var c Contact

//Action - переменная значения выбора
var Action string

//Surname - переменная фамилии
var Surname string

//Find -переменная фамилии для поиска в карте
var Find string

//Len - ...
func (s sortSurname) Len() int {
	return len(s)
}

// Less -...
func (s sortSurname) Less(i, j int) bool {
	return s[i].Surname < s[j].Surname
}

//Swap - ...
func (s sortSurname) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// setContact - добавление контакта в телефонную книгу
func (c *Contact) setContact() {
	fmt.Println("Ввдите Фамилию:")
	fmt.Scanln(&c.Surname)
	fmt.Println("Введите Имя:")
	fmt.Scanln(&c.Name)
	fmt.Println("Введите название организации:")
	fmt.Scanln(&c.Company)
	fmt.Println("Введите телефонный номер:")
	fmt.Scanln(&c.Number)
}

//Options - функция задающая параметры добавления в телефонную книгу, а так же поиска в ней
func Options() {
	fmt.Println("\n\nВыберите действие:\n 1 - поиск контакта\n 2 - добавление нового\n 3 - Показать весь список контактов\n 4 - Выход")
	fmt.Scanln(&Action)
	switch Action {
	case "1":
		FindContact()
	case "2":
		c.setContact()
		TelephoneBook = append(TelephoneBook, c)
		sort.Sort(sortSurname(TelephoneBook))
		Options()
	case "3":
		for k := 0; k < len(TelephoneBook); k++ {
			fmt.Println("\nНомер контакта:", k+1, "\nФамилия:", TelephoneBook[k].Surname, "\nИмя:", TelephoneBook[k].Name, "\nНазвание организации:", TelephoneBook[k].Company, "\nТелефонный номер:", TelephoneBook[k].Number)
		}
		Options()
	case "4":
		fmt.Println("\n\nСпасибо за внимание!")
	default:
		fmt.Println("\n\nНеправильно заданы значения!")
		Options()
	}
}

//FindContact - функция поиска контактов
func FindContact() {
	fmt.Println("Поиск контактов\n\nВведите фамилию:")
	fmt.Scanln(&Find)
	for i := 0; i < len(TelephoneBook); i++ {

		if TelephoneBook[i].Surname == Find {
			fmt.Println("\n\n\n\nФамилия:", TelephoneBook[i].Surname, "\nИмя:", TelephoneBook[i].Name, "\nНазвание организации:", TelephoneBook[i].Company, "\nТелефонный номер:", TelephoneBook[i].Number)
			Options()
		}
	}
	fmt.Println("\n\nКонтакт отсутствует")
	Options()
}

func main() {

	Options()

}
