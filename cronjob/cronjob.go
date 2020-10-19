package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting job;", time.Now())
	// Added time to see output
	time.Sleep(5 * time.Second)
	fmt.Println("Stoping job;", time.Now())
}
