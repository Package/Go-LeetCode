package main

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strings"
)

var (
	maxInt = big.NewInt(math.MaxInt32)
	minInt = big.NewInt(math.MinInt32)
)

func myAtoi(str string) int {
	stripped := strings.TrimLeft(str, " ")
	if len(stripped) == 0 { return 0 } // Edge case

	firstChar := string(stripped[0])

	// Starts with any digit or a "+" or "-"
	startRegex := regexp.MustCompile("^[0-9|\\-|+]")
	digitRegex := regexp.MustCompile("^[0-9]")

	matchesPattern := startRegex.MatchString(stripped)
	if !matchesPattern {
		// Invalid format
		return 0
	}

	// Parse out as many numbers as we can from the string
	parsed := ""
	for index, value := range stripped {
		asString := string(value)

		// Edge case, do not put the "+" or "-" sign into the string
		if index == 0 && ( asString == "+" || asString == "-" ) { continue }

		// Got as far as we can. Now encountered a non-digit element
		if !digitRegex.MatchString(asString) {
			break
		}

		parsed += asString
	}

	// Actually convert
	convertedLarge := big.NewInt(0)
	convertedLarge.SetString(parsed, 10)
	convertedLargeInt := convertedLarge.Int64()

	// If first character was -1 then convert to negative
	if firstChar == "-" {
		convertedLarge = convertedLarge.Mul(convertedLarge, big.NewInt(-1))
		convertedLargeInt *= -1
	}

	// Handle out of bounds of 32bit int
	if convertedLarge.Cmp(minInt) == -1 {
		return math.MinInt32
	}

	if convertedLarge.Cmp(maxInt) == 1 {
		return math.MaxInt32
	}

	return int(convertedLargeInt)
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("20000000000000000000"))
	fmt.Println(myAtoi("   +0 123"))
	fmt.Println(myAtoi("2147483648"))
	fmt.Println(myAtoi("-9223372036854775808"))
}
