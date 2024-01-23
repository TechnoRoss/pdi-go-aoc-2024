package trebuchet

// ExtractNum extracts numbers from an alphanumeric string and calculates their sum
func ExtractNum(input string) int {
	number := extractNumbers(input)
	// fmt.Print("numbers\n")
	//total := sumNumbers(numbers)

	return number
}

func SumDigits(numbers []int) int {
	totals := sumNumbers(numbers)
	// fmt.Print("numbers\n")
	//total := sumNumbers(numbers)

	return totals
}

func GetFirstLast(input string) (string, string) {
	firstNum, lastNum := findFirstLastMatch(input)
	// fmt.Print("numbers\n")
	//total := sumNumbers(numbers)

	return firstNum, lastNum
}
