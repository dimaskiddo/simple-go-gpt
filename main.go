package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	for {
		var question string

		fmt.Println("Question:")

		input := bufio.NewReader(os.Stdin)
		question, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to Get Input from Stdin")
		}

		fmt.Println("")
		fmt.Println("Answer:")

		answer, err := GPTResponse(question)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(answer)

		select {
		case <-sig:
			os.Exit(0)
		case <-time.After(3 * time.Second):
			fmt.Println("")
		}
	}
}
