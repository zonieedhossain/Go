package main

import (
	"fmt"
	"math"
)

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value))/2); i++ {
		if value%i == 0 {

		}
		fmt.Println(i)
		fmt.Println(value)

	}
	fmt.Println("is prime number")
	return value > 1
}

func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {

		}
		fmt.Println(i)
		fmt.Println(value)
	}
	fmt.Println("is prime number")
	return value > 1
}
func main() {
	IsPrimeSqrt(5145)
	IsPrime(5145)

}
