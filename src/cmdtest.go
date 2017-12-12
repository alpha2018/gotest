package main

import (
	"os/exec"
	"fmt"
)

func main() {
	f, err := exec.Command("ls", "go run /Users/leo/Sites/go/src/main.go").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(f))
}