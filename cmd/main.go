package main

import (
	"bufio"
	"calculator/calculator/parser"
	"calculator/calculator/rpn"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Enter your expression")
	fmt.Print(">> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	tokens, err := parser.Parse(line)
	if err != nil {
		log.Fatal(err)
	}

	rpnTokens, ok := rpn.CreateRPN(tokens)
	if !ok {
		fmt.Println("some error")
	}
	res, err := rpn.EvaluateRpn(rpnTokens)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result: ", res)
}
