package main

import ("fmt"
			"math")


type Shape interface{
	Area() float64
}

type Circle struct{
	radius float64	
}

type Rectangle struct{
	width,height float64
}

func (c Circle)Area() float64{
	return math.Pi*c.radius*c.radius
}

func (r Rectangle)Area() float64{
	return r.width*r.height
}

func Print_(s Shape){
		fmt.Printf("Area: %.2f\n",s.Area())
}

func main(){
	c:=Circle{
		radius:5,
	}
	r:=Rectangle{
		width:10,
		height:5,
	}

	Print_(c)
	Print_(r)
}