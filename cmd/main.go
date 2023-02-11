package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"uniq_util/uniq"
)

func parseCmd() (options uniq.Options) {
	flag.BoolVar(&options.CountEntries, "c", false, "подсчитать количество встречаний строки во входных данных")
	flag.BoolVar(&options.OnlyRepeating, "d", false, "вывести только те строки, которые повторились во входных данных.")
	flag.BoolVar(&options.OnlyUnique, "u", false, "вывести только те строки, которые не повторились во входных данных.")
	flag.IntVar(&options.IgnoreFields, "f", 0, "не учитывать первые num_fields полей в строке.")
	flag.IntVar(&options.IgnoreChars, "s", 0, "не учитывать первые num_chars символов в строке")
	flag.BoolVar(&options.IgnoreRegister, "i", false, "не учитывать регистр букв")
	flag.Parse()
	return options
}

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
	opts := parseCmd()
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

	output, err := uniq.Uniq(opts, lines)
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range output {
		fmt.Println(str)
	}
}
