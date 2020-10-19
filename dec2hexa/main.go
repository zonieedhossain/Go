package main

import (
	"fmt"
	"strconv"
)

func main() {
	var decimal int64
	fmt.Print("Enter Decimal Number and driver id:")
	fmt.Scanln(&decimal)
	output := strconv.FormatInt(decimal, 16)
	var s string

	if decimal < 16777215 {
		for i := 0; i < 8-len(output); i++ {
			s = "0" + s
		}
	} else {
		substring := output[1:9]
		fmt.Println("sub ", substring)
	}

	s = s + output
	fmt.Println("Output ", output)
	fmt.Println("S value", s)

}
