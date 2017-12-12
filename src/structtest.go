package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {
	var P person
	P.name = "jack"
	P.age = 11
	fmt.Printf("The Person's name is %s", P.name)
	fmt.Printf("The Person's age is %d", P.age)
}