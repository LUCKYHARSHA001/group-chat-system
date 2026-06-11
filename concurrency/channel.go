package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "Ping from the background!"
	}()

	receivedMsg := <-messages
	fmt.Println(receivedMsg)
}