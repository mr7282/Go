package main

import (
	"fmt"
	"sort"
)

//Создать тип, описывающий контакт в телефонной книге. Создать псевдоним типа телефонной книги (массив контактов) и реализовать для него интерфейс Sort{}.

type contact struct {
	name    string
	surname string
	tel     int
}

type bySurname []contact

func (t bySurname) Len() int {
	return len(t)
}

func (t bySurname) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t bySurname) Less(i, j int) bool {
	return t[i].surname < t[j].surname
}

type byName []contact

func (t byName) Len() int {
	return len(t)
}

func (t byName) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t byName) Less(i, j int) bool {
	return t[i].name < t[j].name
}

func main() {
	myTelephoneBook := []contact{{"Дима", "Грицай", 89265554545}, {"Денис", "Коханов", 89565466565}, {"Дима", "Иванов", 89601232122}, {"Владимир", "Моисеенко", 89162544565}}
	fmt.Println(myTelephoneBook)
	sort.Sort(bySurname(myTelephoneBook))
	fmt.Println(myTelephoneBook)
	sort.Sort(byName(myTelephoneBook))
	fmt.Println(myTelephoneBook)
}
