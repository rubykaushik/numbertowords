package numbertowords

import (
	"errors"
)

var words = [21]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tenwords = [10]string{
	"",
	"",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

// maxNumber denotes the maximum number allowed for Numbertowords function
const maxNumber = 99999

// Convert converts given number to string
func Convert(number int) (string, error) {
	if number < 0 || number > maxNumber {
		return "", errors.New("number not in valid range")
	}

	result := ""

	thousands := number / 1000
	if thousands > 0 {
		result = words[thousands] + " thousand "
		number = number % 1000
		if number == 0 {
			return result, nil
		}
	}

	hundreds := number / 100

	if hundreds > 0 {
		result = result + words[hundreds] + " hundred "
		number = number % 100
		if number == 0 {
			return result, nil
		}
	}

	tens := number / 10
	units := number % 10

	if tens < 2 {
		return result + words[number], nil
	}

	if units == 0 {
		return result + tenwords[tens], nil
	}

	return result + tenwords[tens] + " " + words[units], nil

}
