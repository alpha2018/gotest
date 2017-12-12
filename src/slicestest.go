package main

import "fmt"

func main() {
	primers := []int{2,3,5,7,9}

	var s []int = primers[1:4]
	fmt.Println(s)

	s[2] = 4
	fmt.Println(s)

	var s1 []int
	s1 = append(s1, 1)
	fmt.Println(s1)
}
