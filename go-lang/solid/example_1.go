package main

import "fmt"

type A struct {
	year int
}

func (a A) Greet() { fmt.Println("Hello GolangUK", a.year) }

// B is embedding A here
type B struct {
	A
	name string
}

func (b B) Greet() { fmt.Println("Welcome to GolangUK", b.year, b.name) }

func main() {
	var a A
	a.year = 2016
	var b B
	b.name = "saurabh"
	b.year = 2016
	a.Greet() // Hello GolangUK 2016
	b.Greet() // Welcome to GolangUK 2016
}
