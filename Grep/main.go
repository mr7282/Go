package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func grep(regular, file string) {
	regex, err := regexp.Compile(regular)
	if err != nil {
		return
	}

	fh, err := os.Open(file)
	f := bufio.NewReader(fh)

	if err != nil {
		return
	}
	defer fh.Close()

	buf := make([]byte, 1024)
	for {
		buf, _, err = f.ReadLine()
		if err != nil {
			return
		}

		s := string(buf)
		if regex.MatchString(s) {
			fmt.Printf("%s\n", string(buf))
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 2 {
		grep(flag.Arg(0), flag.Arg(1))
	} else {
		fmt.Printf("Не правильнгое количество аргументов\n")
	}
}
