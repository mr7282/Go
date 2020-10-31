package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type bolid struct {
	brand string
	color string
	speed int
}

func main() {
	var wg sync.WaitGroup
	mercedes := bolid{
		brand: "Mercedes",
		color: "white",
	}

	bmw := bolid{
		brand: "BMW",
		color: "Black",
	}

	toyota := bolid{
		brand: "Toyota",
		color: "Silver",
	}

	bolids := []bolid{
		mercedes,
		bmw,
		toyota,
	}

	for _, car := range bolids {
		wg.Add(1)
		generator := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(1* time.Millisecond)

		car.speed = generator.Intn(150)

		go func(b bolid) {
			defer wg.Done()
			for i := 0; i < b.speed; i++ {
				time.Sleep(1* time.Millisecond)
			}
			fmt.Printf("Finished %v comand, she's time %v\r\n", b.brand, b.speed)
			}(car)
	}
	wg.Wait()
}
