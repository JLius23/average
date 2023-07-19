//average calculates the average of the several numbers. 
package main

import (
	"fmt"
	"example.com/aaverage/datafile"
	"log"
	"github.com/gorilla/mux"
)
func main () {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number 
	} 
	
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
