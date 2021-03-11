package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 0, "Nth digit of Fibonacci sequence")

	flag.Parse()

	fmt.Println("hi", *n+1)
}
