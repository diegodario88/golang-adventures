package main

import "fmt"

func main() {
	var data int

	go func() { data++ }()

	if data > 0 {
		fmt.Println("data has value of ", data)
	} else {
		fmt.Println("data is zero")
	}
}
