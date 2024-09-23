package main

import (
	"fmt"
)

func main() {

	messageLen := 10
	maxmessageLen := 20
	fmt.Println("Trying To send Message :", messageLen, "Maximum Message is :", maxmessageLen)

	// Go Always Make syntax Sorter and make out programming faster

	fmt.Println("******The initial statement of an if block*******")
	//if INITIAL_STATEMENT; CONDITION { }

	//Normal Syntax
	if messageLen <= maxmessageLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message can't Sent")
	}

	//Go Mode Syntax
	if mxmsg := maxmessageLen; messageLen <= mxmsg {
		fmt.Println("Message Sent")
	}

}
