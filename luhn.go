package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

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
	for i, charByte := range number {
		digit, err := strconv.Atoi(string(charByte))
		if err != nil {
			return "", err
		}
		fmt.Printf("%2d: %v - ", i, digit)

		if (len(number)-i)%2 == 0 {
			// Odd digit - add
			fmt.Print("odd, add")
			sum += digit
		} else {
			// Even digit - double and add the digits, then add
			fmt.Print("even, double and add the digits")
			if digit*2 > 9 {
				sum += digit*2 - 9
			} else {
				sum += digit * 2
			}
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
