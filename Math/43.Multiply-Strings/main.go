package main

import (
	"fmt"
)

func main() {
	result := multiply("123", "456")
	fmt.Println(result)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	length1 := len(num1)
	length2 := len(num2)

	result := make([]byte, length1+length2)

	for i := length1 - 1; i >= 0; i-- {
		digit1 := (num1[i] - '0')

		for j := length2 - 1; j >= 0; j-- {
			digit2 := (num2[j] - '0')
			value := digit1 * digit2
			result[i+j+1] += value

			carry := result[i+j+1] / 10 // 進位
			result[i+j] += carry
			result[i+j+1] %= 10
		}
	}

	if result[0] == 0 {
		result = result[1:]
	}

	for i := range result {
		result[i] += '0'
	}
	return string(result)
}
