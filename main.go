package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
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

		questionFirstWord := strings.SplitN(question, " ", 2)[0]
		quitCommand := regexp.MustCompile(strings.ToLower("close|exit|quit|end"))

		if bool(quitCommand.MatchString(strings.ToLower(questionFirstWord))) {
			os.Exit(0)
		}

		fmt.Println("")
		fmt.Println("Answer:")

		err = GPT3Completion(question)
		if err != nil {
			fmt.Println(err.Error())
		}

		select {
		case <-sig:
			os.Exit(0)
		case <-time.After(1 * time.Second):
			fmt.Println("")
		}
	}
}
