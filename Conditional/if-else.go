package main

import (
	"fmt"
)

func main() {

	messageLen := 10
	maxmessageLen := 20
	fmt.Println("Trying To send Message :", messageLen, "Maximum Message is :", maxmessageLen)

	if messageLen <= maxmessageLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message can't Sent")
	}

}
