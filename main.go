package main

import (
	"bufio"
	"fmt"
	"github.com/russmack/nsqscript"
	"os"
	"strings"
)

func main() {
	fmt.Println("\n==========\n nsq repl\n----------")
	fmt.Println("\nWelcome.\n")
	for {
		fmt.Print("\n>")
		in := bufio.NewReader(os.Stdin)
		userText, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Instruction not understood:", err)
			break
		}
		instr := strings.Trim(userText, " \t\r\n")
		if instr == "" {
			continue
		}
		result := nsqscript.ParseLine(instr)
		fmt.Println(result)
	}
}
