package main

import "fmt"

func main(){
	fmt.Println("Welcome to arrays ... ")
	var fruits [4]string
	fruits[0] = "apples"
	fruits[2] = "banana"
	fruits[3] = "ceric"
	fmt.Println("fruits, ",fruits)

	var vegList = [3]string{"potatoes","beans","mushroom"}
	fmt.Println("vegs are here ....", vegList)
}