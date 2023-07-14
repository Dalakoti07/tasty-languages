package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

// todo make implementation a bit robust dont panic everywhere, output info in console

var outputFileName = "outputFile.txt"
var fileName = "../larger.txt"

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

func readFileFromCertainChuck(filePath string, start, end int64,
	interimResults chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	limit := end - start

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

	reader := bufio.NewReader(file)

	if start != 0 {
		_, err = reader.ReadBytes(' ')
		if err == io.EOF {
			fmt.Println("EOF")
			return
		}

		if err != nil {
			panic(err)
		}
	}

	wordHashMap := make(map[string]int)
	var cummulativeSize int64
	for {
		// Break if read size has exceed the chunk size.
		if cummulativeSize > limit {
			break
		}

		b, err := reader.ReadBytes(' ')

		// Break if end of file is encountered.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		cummulativeSize += int64(len(b))
		s := strings.TrimSpace(string(b))
		s = strings.ToLower(s)
		if s != "" {
			// Send the read word in the channel to enter into dictionary.
			if strings.Contains(s, "\n") {
				listy := strings.Split(s, "\n")
				for _, item := range listy {
					item = strings.TrimSpace(item)
					if item != "" {
						wordHashMap[item]++
					}
				}
			} else {
				wordHashMap[s]++
			}
		}
	}
	interimResults <- wordHashMap
}

func main() {
	// calculate the time
	startTime := time.Now()
	totalWordHashMap := make(map[string]int)

	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup

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
