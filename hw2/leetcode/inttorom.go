package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	if num <= 0 || num > 3999 {
		return "Invalid number"
	}

	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result strings.Builder
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			result.WriteString(sym[i])
		}
	}

	return result.String()
}

func main() {
	fmt.Println(intToRoman(3)) 
}
