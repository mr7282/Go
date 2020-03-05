package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Внести в телефонный справочник дополнительные данные. Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных.
func main() {
	telephoneBook := make(map[string]int)
	content, err := ioutil.ReadFile("telephoneBook.json")
	if err != nil {
		fmt.Println("Что то пошло не так при чтении файла!")
	}

	for {
		fmt.Println("Введите 1 если необходимо добавить контакт\nВведите 2 если необходимо показать телефонную книгу целиком\nВведите 3 если необходимо осуществить поиск\nВведите 4 для выхода из телефонной книги")
		var condition string
		fmt.Scanln(&condition)
		switch condition {
		case "1":
			getNewContact(telephoneBook)
		case "2":
			showTelephoneBook(telephoneBook)
		case "3":
			findContact(telephoneBook)
		case "4":
			m, err := json.Marshal(telephoneBook)
			if len(m) != len(content) {
				if err != nil {
					fmt.Println("Что то пошло не так при кодировании в JSON!")
				}
				err = ioutil.WriteFile("telephoneBook.json", m, 0644)
				fmt.Println("Все записи сохранены")
			} else {
				fmt.Println("Без сохранения")
			}
			fmt.Println("Спасибо за пользование наше зписной книжкой")
			break
		default:
			fmt.Println("Неверно введено значение. Попробуй еще раз!")
		}
		if condition == "4" {
			break
		}

	}
}

//getNewContact - this function add new contact in telephonebook
func getNewContact(book map[string]int) {
	fmt.Println("Ввидите Фамилию")
	var surname string
	fmt.Scanln(&surname)
	fmt.Println("Введите телефонный номер")
	var number int
	fmt.Scanln(&number)
	book[surname] = number
}

//findContact - find and shows telephone number related surname
func findContact(book map[string]int) {
	fmt.Println("Введите фамилию для поиска:")
	var surname string
	fmt.Scanln(&surname)
	fmt.Printf("%v\n\n", book[surname])
}

//showTelephoneBook - shows all contact contained telephone book
func showTelephoneBook(book map[string]int) {
	for name, number := range book {
		fmt.Printf("%v - %v\n", name, number)
	}
}
