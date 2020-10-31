package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(mirroredQuery())

	}
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("https://www.techport.ru")
	}()
	go func() {
		responses <- request("https://geekbrains.ru")
	}()
	go func() {
		responses <- request("https://vesna-k.ru")
	}()
	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string {
	r, err := http.Get(hostname)
	if err != nil {
		logrus.WithError(err).Info(("Bad boy"))
	}
	return r.Request.Host
}
