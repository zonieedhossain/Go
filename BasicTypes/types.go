package main

import "fmt"

func main(){
	//bool can either be true or false
	var b bool= true 
	fmt.Printf("b is : '%v'\n", b)
	b = false
	fmt.Printf("b is : '%v'\n", b)
	// Let's try to save bool value = 5
	// b=5 // in this line error given it's not type of bool
	

	// now assign no value that means null just decleard bool value and print
	var b2 bool
	fmt. Printf("None value of b2 is : '%v'\n", b2) // if we print that it's value will be false becasuse of null value
	
}