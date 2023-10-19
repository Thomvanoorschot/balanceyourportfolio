package isinutils

import (
	"fmt"
	"strconv"
	"strings"
)

func CUSIPtoISIN(cusip string, countryCode string) (string, error) {
	cusip = strings.ToUpper(cusip)
	if len(cusip) != 9 {
		return "", fmt.Errorf("Invalid CUSIP: %s", cusip)
	}

	var even []int
	var odd []int
	var checksum int
	var isinDigits string
	for _, char := range countryCode {
		isinDigits += fmt.Sprintf("%d", char-55)
	}
	isinDigits += cusip
	for i, c := range isinDigits {
		digit, _ := strconv.Atoi(string(c))

		if i%2 == 0 {
			even = append(even, digit*2)
		} else {
			odd = append(odd, digit)
		}
	}
	for _, ed := range even {
		for ed > 0 {
			digit := ed % 10
			checksum += digit
			ed /= 10
		}
	}
	for _, od := range odd {
		checksum += od
	}
	checksum = checksum % 10
	checksum = 10 - checksum

	return fmt.Sprintf("%s%s%d", countryCode, cusip, checksum), nil
}
