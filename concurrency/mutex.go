package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var lock sync.Mutex 

func increment() {
	lock.Lock()         
	count++             
	lock.Unlock()       
}

func main() {
	go increment()
	go increment()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Final Count:", count)
}