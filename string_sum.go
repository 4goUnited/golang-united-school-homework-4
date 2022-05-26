package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	// Use when the expression has invalid input characters
	errorInvalidInput = errors.New("There is not valid input characters in your expression")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	input = strings.TrimSpace(input)
	if input == "" {
		return "", fmt.Errorf("Nothing to calculate.\n%w", errorEmptyInput)
	}

	for _, ch := range input {
		if !(unicode.IsDigit(ch) || unicode.IsSpace(ch) || ch == rune('+') || ch == rune('-')) {
			return "", fmt.Errorf("Invalid input.\n%w", errorInvalidInput)
		}
	}

	re := regexp.MustCompile(`-?\d+`)
	sl := re.FindAllString(input, -1)

	if len(sl) != 2 {
		return "", fmt.Errorf("Number of operands must be equal two.\n%w", errorNotTwoOperands)
	}

	firstOperand, err := strconv.Atoi(strings.TrimSpace(sl[0]))
	if err != nil {
		return "", fmt.Errorf("First operand is not valid.\n%w", err)
	}

	secondOperand, err := strconv.Atoi(strings.TrimSpace(sl[1]))
	if err != nil {
		return "", fmt.Errorf("Second operand is not valid.\n%w", err)
	}

	return strconv.Itoa(firstOperand + secondOperand), nil
}
