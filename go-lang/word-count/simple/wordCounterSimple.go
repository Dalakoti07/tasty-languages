package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
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

	// Extract the keys from the hashmap
	var keys []string
	for key := range hashMap {
		keys = append(keys, key)
	}

	// Sort the keys
	sort.Strings(keys)

	// Write the hashmap to the file
	for _, key := range keys {
		value := hashMap[key]
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
	// calculate the time
	startTime := time.Now()

	file, err := os.Open(fileName)
	if err != nil {
		log.Panicf("error in reading file %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Error in closing the file")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	wordHashMap := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		eachWords := strings.Split(line, " ")
		for _, word := range eachWords {
			word = strings.ToLower(word)
			wordHashMap[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Panicf("scanner got error: %v", err)
	}

	// end time
	elapsedTime := time.Since(startTime).Milliseconds()
	println("Time elapsed in milliseconds: ", elapsedTime)

	WriteResultsToFile(wordHashMap)
}
