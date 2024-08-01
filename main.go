package main

import (
	"bufio"
	"calc/evaluator"
	"fmt"
	"os"
	"strings"
)

func main() {
	PrintHelp()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>")
		text, _ := reader.ReadString('\n')
		if strings.HasPrefix(text, "#") {
			switch text {
			case "#exit":
				return
			case "#help":
				PrintHelp()
			}
		} else {
			if result, err := evaluator.Evaluate(text); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		}
	}
}

func PrintHelp() {
	fmt.Println("Simple REPL to evaluate math expresions. For example type: 2*3+4\nSupported operations: +-/*^\n#exit - for exit #help for help")
}
