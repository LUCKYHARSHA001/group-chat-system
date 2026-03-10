package main

import "fmt"

type Book struct{
		name string
		aurthor string
		pages int
}

func (book Book) print_details(){
	fmt.Printf("Book: %s",book.name)
	fmt.Printf("\nAuthor: %s",book.aurthor)
	fmt.Printf("\nPages: %d",book.pages)
}

func main(){
	book1:=Book{
		"A.K",
		"Harsha",
		99,
	}
	book1.print_details()
}