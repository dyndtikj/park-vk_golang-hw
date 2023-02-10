package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"uniq_util/uniq"
)

func readLines(reader io.Reader) ([]string, error) {
	result := make([]string, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return result, err
	}
	return result, nil
}

func main() {
	var opts uniq.Options
	opts.U = true
	opts.I = true
	// debug
	file, err := os.Open("./test_cases/test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	lines, err := readLines(file)
	if err != nil {
		log.Fatal(err)
	}

	output := uniq.Uniq(opts, lines)
	for _, str := range output {
		fmt.Println(str)
	}
}
