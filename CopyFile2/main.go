package main

import (
	"log"
	"os"
)

//Имя копируемого файла и путь передается в качестве первого аргумента
//Имя создаваемого файла и путь передается в качестве второго аргумента
func main() {

	name := os.Args[1]
	newName := os.Args[2]

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

	newFile, err := os.Create(newName)
	if err != nil {
		return
	}

	newFile.Write(fR)

	defer newFile.Close()

}
