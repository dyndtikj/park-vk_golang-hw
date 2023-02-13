package main

import (
	"bufio"
	"calculator/calculator/parser"
	"calculator/calculator/token"
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
	for _, t := range tokens {
		fmt.Println("Token:", t.Literal, t.Type)
		if t.Type == token.OPERATOR {
			fmt.Println("Priority: ", token.Priority[t.Literal])
		}
	}
}
