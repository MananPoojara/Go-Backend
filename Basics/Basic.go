package main;

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("***********************")

	fmt.Println("Variables In Go")

	fmt.Println("var smsSendingLimit \nint smsSendingLimit = 42 \n var costPerSMS float32 \ncostPerSMS = 233.44 \nvar hasPermission bool \nhasPermission = true \nvar username string \nusername = Manan \n'%v' '%.2f' '%v' '%q', smsSendingLimit, costPerSMS, hasPermission, username'")

	fmt.Println("***********************")
	fmt.Println("Another Way to declare variable short way of declaring variable")

	fmt.Println("The walrus operator, :=, declares a new variable and assigns a value to it in one line")
	messageStart := "Hello Dear Sir, my age is"
	age := 19
	messageEnd := "Thank you sir"

	fmt.Println(messageStart, age, messageEnd)

	fmt.Println("***********************")
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)

	fmt.Println("***********************")
	fmt.Println("Generally speaking, languages that compile directly to machine code produce programs that are faster than interpreted programs. \nGo is one of the fastest programming languages, beating JavaScript, Python, and Ruby handily in most benchmarks. \nHowever, Go programs don't run quite as fast as its compiled Rust, Zig, and C counterparts. That said, it compiles much faster than they do, which makes the developer experience super productive. Unfortunately, there are no swordfights on Go teams...")

}
