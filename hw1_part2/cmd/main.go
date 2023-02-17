package main

import (
	"bufio"
	"fmt"
	"homework/hw1_part2/calculator"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	res, err := calculator.Calculate(line)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
