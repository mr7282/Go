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
		fmt.Println(err)
	}
	if content != nil {
		err := json.Unmarshal(content, &telephoneBook)
		if err != nil {
			fmt.Println("Что то пошло не так2")
		}
	}
	for {
		fmt.Println("Введите 1 если необходимо добавить контакт\nВведите 2 если необходимо вывести телефонную книгу целиком\nВведите 3 если необходимо осуществить поиск\nВведите 4 для выхода из телефонной книги")
		var condition string
		fmt.Scanln(&condition)
		switch condition {
		case "1":
			getNewContact(telephoneBook)
		case "2":
			fmt.Println(telephoneBook)
		case "3":
			findContact(telephoneBook)
		case "4":
			fmt.Println("Спасибо за пользование наше зписной ")
			break
		default:
			fmt.Println("Неверно введено значение. Попробуй еще раз!")
		}
		if condition == "4" {
			break
		}

		m, err := json.Marshal(telephoneBook)
		if err != nil {
			fmt.Println("Что то пошло не так3")
		}
		err = ioutil.WriteFile("telephoneBook.json", m, 0644)

	}

	// telephoneBook["Грицай"] = 89261234565
	// telephoneBook["Коханов"] = 89164787489
	// telephoneBook["Иванов"] = 89254541232
	// telephoneBook["Крылов"] = 89035264574
	// telephoneBook["Лазуренко"] = 89265788987

	// m, err := json.Marshal(telephoneBook)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(len(telephoneBook))
	// fmt.Scanln(&telephoneBook)
	// fmt.Println(len(telephoneBook))
	// fmt.Println(string(m))

}

func getNewContact(book map[string]int) {
	fmt.Println("Ввидите Фамилию")
	var surname string
	fmt.Scanln(&surname)
	fmt.Println("Введите телефонный номер")
	var number int
	fmt.Scanln(&number)
	book[surname] = number
}

func findContact(book map[string]int) {
	fmt.Println("Введите фамилию для поиска:")
	var surname string
	fmt.Scanln(&surname)
	fmt.Println(book[surname])
}
