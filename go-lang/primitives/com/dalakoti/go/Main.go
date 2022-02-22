package main

import (
	"fmt"
)

func main() {
	var isAdult bool
	fmt.Printf("value: %v and type: %T \n", isAdult, isAdult)

	var marks int
	fmt.Printf("marks: %v, type: %T \n", marks, marks)

	var nameIs = "My name is saurabh"
	var nameByte = []byte(nameIs)
	fmt.Printf("name in byte %v , type: %T\n", nameByte, nameByte)

	// Rune in Golang are superset of 
}
