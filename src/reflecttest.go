package main

import (
	"fmt"
	//"reflect"
	"log"
	"github.com/go-martini/martini"
)

type Read struct {
	Reader
}

func main()  {
	var x int = 3
	var ptr *int

	ptr = &x
	pointerVariable := *ptr
	test := pointerVariable + 1

	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(test)

	//fmt.Println("type:", reflect.ValueOf(pointerVariable))
}
