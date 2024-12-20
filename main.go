package main

import (
	"fmt"
)

func main() {
	// Measure the performance of the sample function
	duration, allocatedMemory, totalMemory := timer.MeasurePerformance(Radix_in_Memory)

	// Print the results
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("Allocated memory during execution: %d bytes\n", allocatedMemory)
	fmt.Printf("Total memory allocated since start: %d bytes\n", totalMemory)
}