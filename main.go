package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var res int
	if x == 0 {
		return 0
	}
	sign := x / int(math.Abs(float64(x)))
	x = int(math.Abs(float64(x)))
	tens := int(math.Pow(10, math.Floor(math.Log10(float64(x)))))

	for x > 0 {
		rem := int(math.Mod(float64(x), 10))
		res += rem*tens
		x /= 10
		tens/=10
	}

	return res*sign
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return x == reverse(x)
}


func main() {
	fmt.Printf("Is palindrome %v\n", isPalindrome(125))
}