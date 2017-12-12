package main

import (
	"math"
	"fmt"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64{
	return r.width*r.height
}

func (c Circle) area() float64{
	return c.radius*c.radius*math.Pi
}

func main() {
	r1 := Rectangle{2,10}
	r2 := Rectangle{3,19}
	r3 := Circle{7}
	r4 := Circle{3.33}

	fmt.Printf("Area of r1 is: %f",r1.area())
	fmt.Printf("Area of r1 is: %f",r2.area())
	fmt.Printf("Area of r1 is: %f",r3.area())
	fmt.Printf("Area of r1 is: %f",r4.area())

}
