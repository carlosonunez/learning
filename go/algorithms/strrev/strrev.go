package algorithms

import (
	"errors"
)

// This string reversal implementation reverses the string in place.
func reverse(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("string must be of non-zero length.")
	}

	// Adjust the stop point depending on the evenness of the string.
	fulcrum := len(str) - 1/2
	if len(str)%2 == 0 {
		fulcrum = len(str) / 2
	}

	// Do the reversal in place!
	reversedString := str
	for idx := 0; idx < fulcrum; idx++ {
		temp := reversedString[idx]
		reversedString[idx] = reversedString[len(str)-1]
		reversedString[len(str)-1] = temp
	}
	return reversedString, nil
}
