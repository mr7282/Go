package main

import (
	"fmt"
	"log"
	"os"
)

// Имя копируемого файла, создаваемого файла и их пути запрашивает сама программа
func main() {

	var name string
	var newName string

	fmt.Println("Укажите путь и имя копируемого файла")
	fmt.Scanln(&name)

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fR := make([]byte, stat.Size())
	_, err = file.Read(fR)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Укажите новое имя файла, включая путь")
	fmt.Scanln(&newName)
	newFile, err := os.Create(newName)
	if err != nil {
		return
	}

	newFile.Write(fR)

	defer newFile.Close()

}
