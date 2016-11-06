# go-luhn
Luhn checksum example code in Go

## Usage

### As a Library

`go get github.com/dabfleming/go-luhn`

```go
import "github.com/dabfleming/go-luhn"

func main()  {
	// Compute the checksum to complete the number
	checksum, err := luhn.Compute("7992739871")

	// Check if a number is valid
	valid, err := luhn.Check("79927398713")
}
```

### CLI

```
$ go install github.com/dabfleming/go-luhn/cmd/luhn
$ luhn <number to check>
```

#### Example

```
$ luhn 79927398713
79927398713 is valid.
$ luhn 79927398710
79927398710 is not valid.
```

## Notes

- Assumes we are processing a string representing a number so that we can deal with strings of arbitrary length.
- Compute/compute functions return a string containing the checksum digit for easier comparison with the input string in the Check function.
