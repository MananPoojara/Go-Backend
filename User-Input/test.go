package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("******User Input Comma OK******")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Name:")

	//comma ok is way to do try catch in golang
	input, err := reader.ReadString('\n')

	fmt.Println("Welcome,", input)
	fmt.Printf("Welcome %T", input)
	fmt.Println("error gen", err)
}
