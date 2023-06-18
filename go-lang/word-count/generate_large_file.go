package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Open the original file for reading
	originalFile, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// Create the output file
	outputFile, err := os.Create("larger.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Get the size of the original file
	originalFileInfo, err := originalFile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	originalSize := originalFileInfo.Size()

	// Set the desired size of the output file (50 MB)
	desiredSize := int64(50 * 1024 * 1024)

	// Calculate the number of times to replicate the content
	replicationFactor := desiredSize / originalSize

	// Replicate the content of the original file
	for i := int64(0); i < replicationFactor; i++ {
		_, err := io.Copy(outputFile, originalFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Truncate the output file to the desired size
	err = outputFile.Truncate(desiredSize)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File generated successfully!")
}
