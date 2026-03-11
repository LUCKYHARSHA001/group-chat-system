package main

import "fmt"

type Task interface{
	Performtask()
}

type Chef struct{
	name string
	dish string
}

func (c Chef) Performtask() {
	fmt.Printf("Chef %s is cooking %s\n", c.name, c.dish)
}

type Waiter struct{
	name string
}

func (w Waiter)Performtask(){
	fmt.Printf("Waiter %s is serving\n", w.name)
}

func Execute(t Task){
	t.Performtask()
}



func main(){
		c:=Chef{
			name:"Harsha",
			dish:"Chicken",
		}
		w:=Waiter{
			name:"vardhan",
		}
		Execute(c)
		Execute(w)
}