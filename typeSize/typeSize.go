package main

import "fmt"

func main() {
	fmt.Println("**Integers, uints, floats, and complex numbers all have type sizes.**")
	fmt.Println("********")

	fmt.Println("**Positive Whole Numbers** (no decimal) \n'uint' stands for 'unsigned integer'. \nuint uint8 uint16 uint32 uint64 uintptr")
	fmt.Println("********")

	fmt.Println("**Signed Decimal Numbers** \nfloat32 float64 \n**Imaginary Numbers** (rarely used) \ncomplex64 complex128")
	fmt.Println("********")

	accountAgeFloat := 2.6
	accountAgeInt := int64(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
