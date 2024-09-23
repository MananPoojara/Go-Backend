package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//Integer Data Type
	var intVar int = 10
	var int32Var int32 = 32
	var int64Var int32 = 64

	//float Data Type
	var float32Var float32 = 2.44
	var float64Var float64 = 6.32

	// Boolean Data Type
	var Boolean bool = true

	//String Data type
	var str string = "Manan"

	// Print the values
	fmt.Println("Integer Data Types:")
	fmt.Println("intVar:", intVar)
	fmt.Println("int32Var:", int32Var)
	fmt.Println("int64Var:", int64Var)

	fmt.Println("\nFloating Point Data Types:")
	fmt.Println("floatVar:", float32Var)
	fmt.Println("float64Var:", float64Var)

	fmt.Println("\nBoolean Data Type:")
	fmt.Println("boolVar:", Boolean)

	fmt.Println("\nString Data Type:")
	fmt.Println("stringVar:", str)

	// Beaside All Of this there is one more type Called RUNE
	fmt.Println("However, Go also has a special type, rune, which is an alias for int32. This means that a rune is a 32-bit integer, which is large enough to hold any Unicode code point.")

	// When you need to work with individual characters in a string, you should use the rune type. It breaks strings up into their individual characters, which can be more than one byte long.
	// We can use zany characters like emojis and Chinese characters in our strings, and Go will handle them just fine.

	const name = "🐻"
	// if name = "bear" output :constant 'name' byte length: 4 constant 'name' rune length: 4
	// But if name = "🐻" constant 'name' byte length: 4 constant 'name' rune length: 1
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)

}
