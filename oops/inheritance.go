package main 

import "fmt"

type Special struct{
	name string
	experience int
}

func (s Special) Introduction(){
	fmt.Printf("Hi, I am %s ands i have %d's of experience.\n",s.name,s.experience)
}

type Chef struct{
	Special
	starter string
}

func (c Chef) Cook(){
	fmt.Printf("Now %s is making %s\n",c.name,c.starter)
}

func main(){
	myChef:=Chef{
		Special: Special{
			name:"Harsha",
			experience: 3,
		},
		starter:"Chicken",
	}
	myChef.Introduction()
	myChef.Cook()
	fmt.Printf("name: %s",myChef.name)
}