package uniq

import (
	"strconv"
	"strings"
)

type Options struct {
	C bool // подсчитать количество встречаний строки во входных данных
	D bool // вывести только те строки, которые повторились во входных данных.
	U bool // вывести только те строки, которые не повторились во входных данных.
	F int  // не учитывать первые num_fields полей в строке.
	S int  // не учитывать первые num_chars символов в строке
	I bool // не учитывать регистр букв
}

type line struct {
	origin   string
	modified string
}

type repLine struct {
	line  line
	count uint
}

func (r repLine) getOrigin() string {
	return r.line.origin
}

func (r repLine) getModified() string {
	return r.line.modified
}

// функция применяет опции к строке, не меняя ее, возвращает копию
func useOptions(line string, options Options) string {
	// C, D, U - в процессе работы функции применяются
	var result = line
	if options.F > 0 {
		result = strings.Join(strings.Split(result, " ")[options.F:], " ")
	}
	if options.S > 0 {
		result = result[options.S:]
	}
	if options.I {
		result = strings.ToLower(line)
	}
	return result
}

func createLines(options Options, input []string) (result []line) {
	for _, str := range input {
		result = append(result, line{str, useOptions(str, options)})
	}
	return result
}

func findReplicates(input []line) []repLine {
	var dupLines []repLine
	last := -1
	for _, currLine := range input {
		if len(dupLines) != 0 &&
			dupLines[last].getModified() == currLine.modified {
			dupLines[last].count++
		} else {
			dupLines = append(dupLines, repLine{currLine, 1})
			last++
		}
	}
	return dupLines
}

func Uniq(options Options, input []string) []string {
	result := make([]string, 0)
	lines := createLines(options, input)
	repLines := findReplicates(lines)
	for _, repLine := range repLines {
		if options.C {
			result = append(result, strconv.Itoa(int(repLine.count))+" "+repLine.getOrigin())
			continue
		}
		if options.D {
			if repLine.count > 1 {
				result = append(result, repLine.getOrigin())
			}
			continue
		}
		if options.U {
			if repLine.count == 1 {
				result = append(result, repLine.getOrigin())
			}
		}
	}
	return result
}
