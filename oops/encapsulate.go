package main

import "fmt"

type Cook struct{
	dish string
	chef string
}

func (c Cook)Print_(){
	fmt.Printf("chef %s is making %s\n",c.chef,c.dish)
}

func main(){
	myCook:=Cook{
		chef: "Harsha",
		dish:"pastha",
	}
	myCook.Print_()
}