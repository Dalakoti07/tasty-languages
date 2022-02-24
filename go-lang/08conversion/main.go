package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to pizza store"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for giving rating, ", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err!=nil{
		fmt.Println("error is", err)
	}else{
		fmt.Printf("final rating: %v \n", numRating+1)
	}
}
