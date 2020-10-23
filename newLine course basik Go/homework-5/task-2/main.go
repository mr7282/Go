package main

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
)

func main() {
	innnerPath := os.Args[1]
	innerGrep := os.Args[2]
	// var list []string
	openDir, err := os.Open(innnerPath)
	if err != nil {
		log.Fatal(err)
	}
	defer openDir.Close()

	filepath.Walk(innnerPath, func(path string, info os.FileInfo, err error) error {
		outinfo := info.a
		fmt.Println()
		return nil
	})


	// readDir, err := openDir.Readdir(-1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, rd := range readDir {
	// 	list = append(list, rd.Name())
	// }
	fmt.Println(innerGrep)
	// fmt.Println(list)

}