package luhn

import (
	"fmt"
	"strings"
)

var evenDigits = []int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// Compute computes the Luhn checksum of the number provided.
func Compute(number string) (string, error) {
	if verifyNumeric(number) == false {
		return "", fmt.Errorf("Error, input string '%v' not numeric.", number)
	}

	fmt.Printf("Compute Input: %v, len(%v)\n", number, len(number))

	return compute(number)
}

// Check checks if the number is valid by the Luhn algorithm.
func Check(number string) (bool, error) {
	if verifyNumeric(number) == false {
		return false, fmt.Errorf("Error, input string '%v' not numeric.", number)
	}
	checksum, err := compute(number[0 : len(number)-1])
	if err != nil {
		return false, err
	}
	fmt.Printf("Check: Returned checksum %#v, Input checksum %#v\n", checksum, number[len(number)-1:])
	return checksum == number[len(number)-1:], nil
}

func compute(number string) (string, error) {
	sum := 0
	mod := len(number) % 2
	for i, charByte := range number {
		fmt.Printf("%2d: %v - ", i, string(charByte))

		if i%2 == mod {
			// Odd digit - add
			fmt.Print("odd, add")
			sum += int(charByte - '0')
		} else {
			// Even digit - double and add the digits, then add
			fmt.Print("even, double and add the digits")
			sum += evenDigits[charByte-'0']
		}
		fmt.Println()
	}
	fmt.Printf("Sum is %v\n", sum)
	checkDigit := fmt.Sprintf("%d", (10-(sum%10))%10)
	return checkDigit, nil
}

func verifyNumeric(s string) bool {
	if strings.Trim(s, "0123456789") == "" {
		return true
	}
	return false
}
