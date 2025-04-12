package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Vai executar infinitamente (não tem garbage collection para goroutine)
// só vai finalizar quando a main terminar
func sender(payload string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", payload, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
	}
}

func main() {
	c := make(chan string)

	go sender("01JRNGRTNQFPQCN20HGSYRZSMR", c)

	for range 5 {
		fmt.Printf(
			"Sender sends: %q\n ",
			<-c, // Vai bloquear a main até chegar um valor no canal
		)
	}

	fmt.Println("Bye Bye ...")
}
