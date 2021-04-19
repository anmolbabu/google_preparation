package main

import (
	"fmt"
)

type Digit struct {
	Value int
	Index int
}

func toString(arr []int) string {
	resultStr := ""

	for _, ele := range arr {
		if (ele == 0) && (resultStr == "") {
			continue
		} else {
			resultStr = fmt.Sprintf("%s%d", resultStr, ele)
		}
	}

	if resultStr == "" {
		return "0"
	}

	return resultStr
}

func multiply(num1 string, num2 string) string {
	result := make([]int, len(num1) + len(num2))

	for num1Idx := len(num1) - 1; num1Idx >= 0; num1Idx-- {
		for num2Idx := len(num2) - 1; num2Idx >= 0; num2Idx-- {
			num1Digit := int(num1[num1Idx] - '0')
			num2Digit := int(num2[num2Idx] - '0')

			digitsPro := num1Digit * num2Digit
			digitsProLSB := digitsPro % 10
			digitsProMSB := digitsPro / 10

			digitProLSBIdx := num1Idx + num2Idx + 1
			digitProMSBIdx := digitProLSBIdx - 1

			addDigitProductToRes(&result, Digit{Value: digitsProLSB, Index: digitProLSBIdx}, Digit{Value: digitsProMSB, Index: digitProMSBIdx})
		}
	}

	return toString(result)
}

func addDigitProductToRes(result *[]int, lsbDigit Digit, msbDigit Digit) {
	lsbSum := (*result)[lsbDigit.Index] + lsbDigit.Value
	lsbCarry := lsbSum / 10
	lsbSum = lsbSum % 10

	msbSum := (*result)[msbDigit.Index] + msbDigit.Value + lsbCarry
	msbCarry := msbSum/10
	msbSum = msbSum % 10

	(*result)[lsbDigit.Index] = lsbSum
	(*result)[msbDigit.Index] = msbSum
	if msbCarry > 0 {
		(*result)[msbDigit.Index-1] = (*result)[msbDigit.Index-1] + msbCarry
	}
}

func main() {
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
	fmt.Println(multiply("0", "0"))
	fmt.Println(multiply("0000", "0000"))
}
