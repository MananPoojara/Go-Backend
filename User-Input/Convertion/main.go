package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome To Dominos")

	// Make Reader
	reader := bufio.NewReader(os.Stdin)

	// now take data from reader and ReadString('\n') stop when we enter
	fmt.Println("Enter Rating Between 1 To 5")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks For,", input)
	//here we do convertion
	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	//error : strconv.ParseFloat: parsing "4\n": invalid syntax
	// we are storing \n also so what to do we can trim is hmmmmmmmmmm

	if err != nil { // error is not nil/null
		fmt.Println(err)
	} else {
		fmt.Println("We Add 1 Rating from our side", rating+1)
	}

}
