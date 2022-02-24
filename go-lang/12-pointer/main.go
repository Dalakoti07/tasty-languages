package main

import "fmt"

func main(){
	myNumber := 7
	fmt.Println("pointer .... ")
	
	ptr  := &myNumber 

	fmt.Println("pointer pointing to ", ptr)
	fmt.Println("pointer value is ", *ptr)
	*ptr = *ptr * 2
	fmt.Println("pointer value is ", myNumber)
}