package trebuchet

import (
	"regexp"
	"strconv"
	"strings"
)

func getFirstLast(digits []int) (int, int) {
	first := digits[0]
	last := digits[len(digits)-1]

	return first, last
}

func extractNumbers(input string) int {
	// Define a regular expression pattern to match numbers
	regex_digit := regexp.MustCompile(`\d+`)

	// Find all matches in the input string
	matches := regex_digit.FindAllString(input, -1)

	// Concatenate the matches into a single string
	concatenatedDigits := ""

	for _, match := range matches {
		concatenatedDigits += match
	}

	strings.Join([]string{concatenatedDigits}, "")

	// Handle single digit
	if len(concatenatedDigits) == 1 {
		var sb strings.Builder
		sb.WriteString(concatenatedDigits)
		sb.WriteString(concatenatedDigits)
		concatenatedDigits = sb.String()
		// doubleDigit := fmt.Sprintf("%s%s", string(concatenatedDigits[0]), string(concatenatedDigits[0]))
		// concatenatedDigits := string(doubleDigit)
	}

	// get first and last digit
	firstDigit := concatenatedDigits[0]
	lastDigit := concatenatedDigits[len(concatenatedDigits)-1]

	var sb2 strings.Builder
	sb2.WriteString(string(firstDigit))
	sb2.WriteString(string(lastDigit))
	concatenatedDigits = sb2.String()
	//print(concatenatedDigits)

	// Convert the concatenated digits to an integer
	num, err := strconv.Atoi(concatenatedDigits)
	if err != nil {
		return 0 // Return 0 if there was an error
	}

	return num
}

func findFirstLastMatch(input string) (string, string) {
	// Define regular expressions for numbers and words
	wordRegex := regexp.MustCompile(`(?:one|two|three|four|five|six|seven|eight|nine)+`)
	//digitRegex := regexp.MustCompile(`\d+`)

	// Find all matches for words and digits in the input
	wordMatches := wordRegex.FindAllString(input, -1)
	//digitMatches := digitRegex.FindAllString(input, -1)

	// Initialize variables to store the first and last matches
	firstMatch := ""
	lastMatch := ""

	// Check if there are any word matches
	if len(wordMatches) > 0 {
		firstMatch = wordMatches[0]
		lastMatch = wordMatches[len(wordMatches)-1]
	}

	print(firstMatch, lastMatch)
	// If there are no word matches or the last digit comes after the last word match, consider digits
	// if len(digitMatches) > 0 && (lastMatch == "" || digitRegex.FindStringIndex(input)[1] > wordRegex.FindStringIndex(input)[1]) {
	// 	if firstMatch == "" {
	// 		firstMatch = digitMatches[0]
	// 	}
	// 	lastMatch = digitMatches[len(digitMatches)-1]
	// }

	return firstMatch, lastMatch
}
