package main

import (
	"bufio"
	"calculator/calculator"
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
	res, err := calculator.Calculate(line)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result: ", res)
}
