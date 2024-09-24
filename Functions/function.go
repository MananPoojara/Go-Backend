package main

import (
	"fmt"
)

// In Go Function Syntax
// func sum(x1 int, x2 int) int{} it's called function signature and returning int taking two int numbers

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func main() {
	test("Hello", "world")
}
