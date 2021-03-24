package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 0, "Nth digit of Fibonacci sequence")

	flag.Parse()

	mem := make([]int, *n+1)

	fmt.Println(fibDB(*n, mem))

}

func fibDB(n int, mem []int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	} else if mem[n] != 0 {
		return mem[n]
	}

	mem[n] = fibDB(n-1, mem) + fibDB(n-2, mem)

	return mem[n]
}
