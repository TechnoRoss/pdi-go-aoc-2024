package trebuchet

func sumNumbers(numbers []int) int {
	total := 0

	// Iterate over the input strings
	for _, num := range numbers {
		total += num
	}

	return total
}
