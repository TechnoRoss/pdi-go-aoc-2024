package main

import (
	"bufio"
	"fmt"
	"os"
	"trebuchet"
)

func main() {
	//input_file := "inputs/day1/testinput.txt"
	input_file := "inputs/day1/testinput2.txt"
	//input_file := "inputs/day1/input.txt"
	// Open the file for reading
	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// Initialize a slice to store the extracted numbers
	extractedNumbers := []int{}

	// Read lines from the file and process each line
	for scanner.Scan() {
		line_in := scanner.Text()

		// Extract numbers from the line using the module
		//num_extracted := trebuchet.ExtractNum(line_in)

		firstMatch, lastMatch := trebuchet.GetFirstLast(line_in)
		print(firstMatch, lastMatch)

		// Append the line's number to the slice
		//extractedNumbers = append(extractedNumbers, num_extracted)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Calculate the total sum of all extracted numbers
	totalSum := trebuchet.SumDigits(extractedNumbers)

	//Print the total sum
	print("Total: ", totalSum)

}
