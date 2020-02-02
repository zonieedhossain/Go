package main

import (
	"fmt"
	"math"
)

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func SieveOfEratosthenes(value int) {
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if f[i] == false {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println("")
}

func main() {
	for i := 1; i <= 100; i++ {
		if IsPrime(i) {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println("")
	for i := 1; i <= 100; i++ {
		if IsPrimeSqrt(i) {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println("")
	SieveOfEratosthenes(100)
}
