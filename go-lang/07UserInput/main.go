package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for giving rating, ", input)
	fmt.Printf("input type: %T \n", input)
}