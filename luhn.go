package luhn

import (
	"fmt"
	"strings"
)

// Pre-compute the resulting sum for the even digits. (Multiply by 2, add the digits.)
var evenDigits = []int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// Compute computes the Luhn checksum of the number provided.
func Compute(number string) (string, error) {
	if verifyNumeric(number) == false {
		return "", fmt.Errorf("Error, input string '%v' not numeric.", number)
	}

	return compute(number), nil
}

// Check checks if the number is valid by the Luhn algorithm.
func Check(number string) (bool, error) {
	if verifyNumeric(number) == false {
		return false, fmt.Errorf("Error, input string '%v' not numeric.", number)
	}

	checksum := compute(number[0 : len(number)-1])
	return checksum == number[len(number)-1:], nil
}

func compute(number string) string {
	sum := 0
	mod := len(number) % 2

	for i, charByte := range number {
		if i%2 == mod {
			// Odd digit - add value
			sum += int(charByte - '0')
		} else {
			// Even digit - add pre-computed value
			sum += evenDigits[charByte-'0']
		}
	}

	checkDigit := fmt.Sprintf("%d", (10-(sum%10))%10)
	return checkDigit
}

func verifyNumeric(s string) bool {
	if strings.Trim(s, "0123456789") == "" {
		return true
	}
	return false
}
