package main

import (
	"bufio"
	"fmt"
	"homework/hw1_part2/calculator"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Printf("failed to scan %s\n", err)
		return
	}
	res, err := calculator.Calculate(line)
	if err != nil {
		fmt.Printf("failed to use calc utility %s\n", err)
		return
	}
	fmt.Println(res)
}
