package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	c := bufio.NewScanner(os.Stdin)
	for c.Scan() {
		fmt.Println("Hello, World.")
		fmt.Println(c.Text())
	}
	if err := c.Err(); err != nil {
		fmt.Println(err)
	}
}
