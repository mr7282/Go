package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)


func main() {
	Tick := make(<-chan time.Time)
	Tick = time.Tick(1 * time.Second)
	cancel := make(chan int)
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		ex := ""
		for {
			fmt.Scanln(&ex)
			if ex == "exit" {
				fmt.Println("chanle is ok")
				cancel <- 1
			}
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, cancel, Tick)
		if <-cancel == 1 { // Можно так
			return
		}
	}
}

func handleConn(c net.Conn, cancel chan int, tick <-chan time.Time) {
	defer c.Close()
	for {
		// select {                         // Или так - ждет нового соединения
		// case <-cancel:
		// 		fmt.Println("Server is down")
		// 		// return
		// case <-tick:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\r"))
			if err != nil {
				return
			}
		// }
	}
}