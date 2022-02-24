package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("welcome to slices")

	var fruitsList = []string{"Apples","Bananas"}
	fmt.Println("fruitsList: ",fruitsList)
	fruitsList = append(fruitsList, "lichi")
	fmt.Println(fruitsList)

	fmt.Println("appending: ",fruitsList[1:])

	// another way
	highScores := make([]int, 4)
	highScores[0] = 1
	highScores[1] = 2
	highScores[2] = 11
	highScores[3] = 4
	fmt.Println("highScore", highScores)
	highScores = append(highScores, 5, 6, 7)
	fmt.Println("highScore", highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println("is sorted? :", sort.IntsAreSorted(highScores))

}