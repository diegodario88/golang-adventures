package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sender(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := range 10 {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
		}

		close(c) // Sender deve ser o responsável for fechar o canal se não deadlock
	}()

	return c
}

func main() {
	svcBlue := sender("blue")
	svcBrown := sender("brown")

	for msg := range svcBlue {
		fmt.Println(msg)
	}

	for msg := range svcBrown {
		fmt.Println(msg)
	}

	fmt.Println("Bye Bye ...")
}
