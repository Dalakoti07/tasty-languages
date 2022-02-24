package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to study of time")
	currentTime := time.Now()
	fmt.Printf("current time %v", currentTime)

	fmt.Printf("\nformatted time, %v\n", currentTime.Format("01-02-2006 Monday"))
}