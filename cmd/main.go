package main

import (
	"bufio"
	"calculator/calc_util"
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
	tokens, err := calc_util.Parse(line)
	if err != nil {
		log.Fatal(err)
	}
	for _, token := range tokens {
		fmt.Println("Token:", token.Literal)
		if token.Type == calc_util.OPERATOR {
			fmt.Println("Priority: ", calc_util.Priority[token.Literal])
		}
	}
}
