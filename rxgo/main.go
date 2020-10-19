package main

import (
	"fmt"

	"github.com/reactivex/rxgo"
)

func main() {
	observable := rxgo.Just("Hello, World!")()
	ch := observable.Observe()
	item := <-ch
	fmt.Println(item.V)
}
