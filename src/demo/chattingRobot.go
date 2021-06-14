package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	//read data from stdin
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")
	//read util there is a '\n' in the end
	input, err := inputReader.ReadString('\n')
	if err != nil{
		fmt.Printf("An error cccured: %s\n", err)
	} else {
		//delete '\n' with slice
		if len(input) >= 2 {
			input = input[:len(input)-2]
		}

		fmt.Printf("Hello, %s! What can I do for you?\n", input)
	}

	for{
		input, err := inputReader.ReadString('\n')
		if err != nil{
			fmt.Printf("An error cccured: %s\n", err)
		}
		if len(input) >= 2 {
			input = input[:len(input)-1]
		}

		//convert input data to lower case
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			//exit normally
			os.Exit(0)
		default:
			fmt.Println("Sorry, I didn't catch you.")
		}
	}
}
