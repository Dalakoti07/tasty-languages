package main

import ("fmt"
	"strconv"
)

var languageName = "go-lang"

var (
	name = "saurabh"
	age = 21
	langauge = "go lang"
)

func main() {
	var i float32 = 42
	fmt.Println(i)

	var strInt = strconv.Itoa(int(i))

	fmt.Printf("%v , %T",i, i)

	fmt.Printf("\n%v , %T",strInt, strInt)
	
	fmt.Println("\nhi again")
}