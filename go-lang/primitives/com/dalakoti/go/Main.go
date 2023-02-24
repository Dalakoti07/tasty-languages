package main

import (
	"fmt"
)

func main() {
	// means default value is false
	var isAdult bool
	fmt.Printf("value: %v and type: %T \n", isAdult, isAdult)

	// means default value is 0
	var marks int
	fmt.Printf("marks: %v, type: %T \n", marks, marks)

	var nameIs = "My name is saurabh"
	var nameByte = []byte(nameIs)
	fmt.Printf("name in byte %v , type: %T\n", nameByte, nameByte)

	// Rune in Golang are superset of
}
