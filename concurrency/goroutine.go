package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from a background goroutine!")
}

func main() {

	go sayHello() 

	fmt.Println("Hello from the main program!")

	time.Sleep(100 * time.Millisecond) 
}