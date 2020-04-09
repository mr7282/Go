package main

import (
	"flag"
	"log"
	"os"
)

//Задание номер 4 - Напишите утилиту для копирования файлов, используя пакет flag.

func main() {
	userFlag := flag.String("f", "", "the file from which")
	flag.Parse()
	from, err := os.Open(*userFlag + ".txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer from.Close()

	fromStat, err := from.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	fromByte := make([]byte, fromStat.Size())
	from.Read(fromByte)

	newFileCreate, err := os.Create(*userFlag + ".copy.txt")
	if err != nil {
		log.Fatalln(err)
	}
	newFileCreate.Write(fromByte)
}
