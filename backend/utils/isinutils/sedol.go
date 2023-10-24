package isinutils

import (
	"fmt"
	"strconv"
	"strings"
)

func SEDOLtoISIN(sedol string, countryCode string) (string, error) {
	sedol = strings.ToUpper(sedol)
	if len(sedol) != 7 {
		return "", fmt.Errorf("Invalid SEDOL: %s", sedol)
	}

	var even []int
	var odd []int
	var checksum int
	var isinDigits string
	for _, char := range countryCode {
		isinDigits += fmt.Sprintf("%d", char-55)
	}
	isinDigits += "00"
	isinDigits += sedol
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

	return fmt.Sprintf("%s00%s%d", countryCode, sedol, checksum), nil
}
