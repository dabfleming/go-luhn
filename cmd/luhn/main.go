package main

import (
	"fmt"
	"os"

	"github.com/dabfleming/go-luhn"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage: luhn <number>\n")
	}

	number := args[1]

	valid, err := luhn.Check(number)
	if err != nil {
		fmt.Printf("Encountered an error checking the number: %v\n", err)
		return
	}

	if valid {
		fmt.Printf("%v is valid.\n", number)
	} else {
		fmt.Printf("%v is not valid.\n", number)
	}
}
