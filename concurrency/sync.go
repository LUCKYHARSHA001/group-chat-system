package main

import (
	"fmt"
	"time"
)

func main(){
	go count("Harsha")
	count("Vardhan")
}

func count(thing string){
	for i:=0;true;i++{
		fmt.Println(i,thing)
		time.Sleep(time.Millisecond*500)
	}
}