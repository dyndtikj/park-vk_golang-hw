package main

import (
	"bufio"
	"flag"
	"fmt"
	"homework/hw1_part1/uniq"
	"io"
	"log"
	"os"
)

const correctOpts = "go run main.go [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]"

type ioSettings struct {
	input  string
	output string
}

func parseCmd() (options uniq.Options, ioSet ioSettings) {
	flag.BoolVar(&options.CountEntries, "c", false, "подсчитать количество встречаний строки во входных данных")
	flag.BoolVar(&options.OnlyRepeating, "d", false, "вывести только те строки, которые повторились во входных данных.")
	flag.BoolVar(&options.OnlyUnique, "u", false, "вывести только те строки, которые не повторились во входных данных.")
	flag.IntVar(&options.IgnoreFields, "f", 0, "не учитывать первые num_fields полей в строке.")
	flag.IntVar(&options.IgnoreChars, "s", 0, "не учитывать первые num_chars символов в строке")
	flag.BoolVar(&options.IgnoreRegister, "i", false, "не учитывать регистр букв")
	flag.Parse()
	if flag.NArg() > 0 {
		ioSet.input = flag.Arg(0)
	}
	if flag.NArg() > 1 {
		ioSet.output = flag.Arg(1)
	}
	return
}

func readData(settings ioSettings) (lines []string, err error) {
	file := os.Stdin
	if len(settings.input) > 0 {
		file, err := os.Open(settings.input)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %w", err)
		}
		defer func(file *os.File) {
			e := file.Close()
			if err == nil {
				err = fmt.Errorf("failed to close file %w", e)
			}
		}(file)
	}
	lines, err = readLines(file)
	if err != nil {
		return nil, fmt.Errorf("failed read lines from file %w", err)
	}
	return lines, err
}

func readLines(reader io.Reader) ([]string, error) {
	var result []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan lines %w", err)
	}
	return result, nil
}

func writeData(settings ioSettings, data []string) (err error) {
	file := os.Stdout
	if len(settings.output) > 0 {
		file, err = os.Create(settings.output)
		if err != nil {
			return fmt.Errorf("failed to create file %w", err)
		}
		defer func(file *os.File) {
			if err = file.Close(); err != nil {
				log.Fatal(err)
			}
		}(file)
	}
	for _, str := range data {
		if _, err = file.WriteString(str + "\n"); err != nil {
			return
		}
	}
	return nil
}

func main() {
	opts, settings := parseCmd()
	if !opts.IsValid() {
		fmt.Printf("invalid arguments, usage:\n%s", correctOpts)
		return
	}

	lines, err := readData(settings)
	if err != nil {
		fmt.Printf("failed to read data %s\n", err)
		return
	}

	output, err := uniq.Uniq(opts, lines)
	if err != nil {
		fmt.Printf("failed to use uniq utility %s\n", err)
		return
	}

	if err = writeData(settings, output); err != nil {
		fmt.Printf("failed to write data %s\n", err)
		return
	}
	return
}
