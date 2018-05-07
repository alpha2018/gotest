package main

import "fmt"

func reverse(p *[]int) {
	for i, j := 0, len(*p)-1; i < len(*p)/2; i, j = i+1, j-1 {
		fmt.Println(i,j,len(*p)-1,p)
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func main() {
	v := []int{1, 2, 3, 4}
	reverse(&v)
	fmt.Println(v)
}