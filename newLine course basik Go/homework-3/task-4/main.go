package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
)

//Внести в телефонный справочник дополнительные данные. Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных.
func main() {
	telephoneBook := make(map[string]int)

	telephoneBook["Грицай"] = 89261234565
	telephoneBook["Коханов"] = 89164787489
	telephoneBook["Иванов"] = 89254541232
	telephoneBook["Крылов"] = 89035264574
	telephoneBook["Лазуренко"] = 89265788987

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
