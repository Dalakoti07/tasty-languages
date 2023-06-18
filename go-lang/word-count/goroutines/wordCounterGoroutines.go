package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

// todo make implementation a bit robust dont panic everywhere, output info in console

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

func readFileFromCertainChuck(filePath string, start, end int64,
	interimResults chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		log.Panicf("error in reading file %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Error in closing the file")
		}
	}(file)

	// set the position in file to start position, seek from file start hence flag as 0
	_, err = file.Seek(start, 0)
	if err != nil {
		log.Panicf("Error seeking file: %f", err)
	}

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
	fmt.Printf("sending result for %v-%v \n", start, end)
	interimResults <- wordHashMap

	if err := scanner.Err(); err != nil {
		log.Panicf("scanner got error: %v", err)
	}
}

func main() {
	// calculate the time
	startTime := time.Now()
	totalWordHashMap := make(map[string]int)

	numWorkers := 5
	var wg sync.WaitGroup

	// todo improve this concurrent design
	// create channel to receive hashmap from channel, and we would be using fan in approach
	interimResulFromWorkers := make(chan map[string]int)

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		panic("Error getting file info")
	}
	fileSize := fileInfo.Size()
	chuckSize := fileSize / int64(numWorkers)

	fmt.Printf("\nfile info: size %v\n", fileSize)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		start := int64(i) * chuckSize
		var end int64
		if i == numWorkers-1 {
			end = fileSize
		} else {
			end = start + chuckSize
		}
		go readFileFromCertainChuck(fileName, start, end, interimResulFromWorkers, &wg)
	}

	// create goroutine to close channel when all workers are finish
	go func() {
		wg.Wait()
		// end time
		elapsedTime := time.Since(startTime).Milliseconds()
		println("Time elapsed in milliseconds: ", elapsedTime)
		WriteResultsToFile(totalWordHashMap)

		close(interimResulFromWorkers)
	}()

	for result := range interimResulFromWorkers {
		for key, value := range result {
			totalWordHashMap[key] += value
		}
	}

}
