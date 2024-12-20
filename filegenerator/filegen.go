package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	numFiles = 6
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the number of lines for each file
	lineCounts := []int{
		1000,         // 1st file
		10000,        // 2nd file
		100000,       // 3rd file
		10000000,     // 4th file
		1000000000,   // 5th file
		100000000000, // 6th file
	}

	// Generate the files
	for i, lines := range lineCounts {
		fileName := fmt.Sprintf("test_file_%d.txt", i+1)
		err := generateFile(fileName, lines)
		if err != nil {
			fmt.Printf("Error generating file %s: %v\n", fileName, err)
		} else {
			fmt.Printf("Successfully generated file %s with %d lines.\n", fileName, lines)
		}
	}
}

// generateFile generates a file with the specified number of lines, each containing a random six-digit integer.
func generateFile(fileName string, lines int) error {
	// Create the file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write random six-digit integers to the file
	for i := 0; i < lines; i++ {
		randomInt := rand.Intn(900000) + 100000 // Generate a random six-digit integer
		_, err := fmt.Fprintln(file, randomInt)
		if err != nil {
			return err
		}
	}

	return nil
}