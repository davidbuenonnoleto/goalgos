package main

import "fmt"

func main() {

	num := 8

	list := [6]int{2, 3, 1, 7, 8, 9}

	for i := range list {
		if list[i] == num {
			fmt.Println("found", list[i])
		}
	}
}
