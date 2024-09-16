package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func isValidCardNumber(card string) bool {
	// Ensure the card number contains exactly 16 digits
	match, _ := regexp.MatchString(`^[456]\d{15}$`, card)
	if !match {
		return false
	}

	// Check for consecutive digits repeating more than 3 times
	for i := 0; i < len(card)-3; i++ {
		if card[i] == card[i+1] && card[i+1] == card[i+2] && card[i+2] == card[i+3] {
			return false
		}
	}

	// Luhn algorithm implementation
	return checkLuhn(card)
}

func checkLuhn(card string) bool {
	sum := 0
	double := false

	for i := len(card) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(card[i]))

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}

	return sum%10 == 0
}

func main() {
	card := "4567891234567890"
	if isValidCardNumber(card) {
		fmt.Println("Valid Card")
	} else {
		fmt.Println("Invalid Card")
	}
}
