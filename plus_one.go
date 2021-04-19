package main

func plusOne(digits []int) []int {
	digitIdx := len(digits) - 1
	carry := 1

	for digitIdx >= 0 {
		currDigit := digits[digitIdx] + carry
		carry = currDigit/10
		currDigit = currDigit % 10
		digits[digitIdx] = currDigit

		digitIdx--
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}

