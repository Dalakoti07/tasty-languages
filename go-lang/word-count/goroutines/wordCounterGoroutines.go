package main

import (
	"fmt"
	"os"
)

var outputFileName = "outputFile.txt"
var fileName = "../sample.txt"

func WriteResultsToFile(hashMap map[string]int) {
	file, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the hashmap to the file
	for key, value := range hashMap {
		line := fmt.Sprintf("%s: %v\n", key, value)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Hashmap written to output.txt")
}

func main() {

}
