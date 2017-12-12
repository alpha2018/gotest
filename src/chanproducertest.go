package main

import "time"

func main() {
	producer()
}

func producer(c chan int64, max int) {
	defer
		close(c)

	for i:= 0; i < max; i ++ {
		c <- time.Now().Unix()
	}
}
