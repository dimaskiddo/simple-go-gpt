package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var question string

	fmt.Println("Question:")

	input := bufio.NewReader(os.Stdin)
	question, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to Get Input from Stdin")
	}

	fmt.Println("")
	fmt.Println("Answer:")

	answer, err := ChatGPTResponse(question)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(answer)
}
