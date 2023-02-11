package uniq

import (
	"errors"
	"strconv"
	"strings"
)

type Options struct {
	CountEntries   bool // подсчитать количество встречаний строки во входных данных
	OnlyRepeating  bool // вывести только те строки, которые повторились во входных данных
	OnlyUnique     bool // вывести только те строки, которые не повторились во входных данных
	IgnoreFields   int  // не учитывать первые num_fields полей в строке
	IgnoreChars    int  // не учитывать первые num_chars символов в строке
	IgnoreRegister bool // не учитывать регистр букв
}

// структура описывает одну строку, необходима для дальнейшей обрпботки
type line struct {
	origin   string // сама оригинальная строка
	modified string // строка после применения опций
}

type repLine struct {
	line  line //информация о строке (оригнальная и после применения опций)
	count uint //количество данных строк
}

func (r repLine) getOrigin() string {
	return r.line.origin
}

func (r repLine) getModified() string {
	return r.line.modified
}

func (r repLine) isUniq() bool {
	return r.count == 1
}

func (o Options) IsValid() bool {
	//вспомнили дискретку
	return o.IgnoreChars >= 0 && o.IgnoreFields >= 0 &&
		(!o.OnlyUnique && !o.OnlyRepeating ||
			(!o.CountEntries && ((o.OnlyRepeating && !o.OnlyUnique) || (!o.OnlyRepeating && o.OnlyUnique))))
}

// Функция применяет опции к строке, не меняя ее, возвращает копию
func useOptions(line string, options Options) string {
	var result = line
	if options.IgnoreFields > 0 {
		result = strings.Join(strings.Split(result, " ")[options.IgnoreFields:], " ")
	}
	if options.IgnoreChars > 0 {
		result = result[options.IgnoreChars:]
	}
	if options.IgnoreRegister {
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

func Uniq(options Options, input []string) ([]string, error) {
	result := make([]string, 0)
	if !options.IsValid() {
		return result, errors.New("invalid arguments")
	}
	lines := createLines(options, input)
	repLines := findReplicates(lines)
	for _, repLine := range repLines {
		if options.CountEntries {
			result = append(result, strconv.Itoa(int(repLine.count))+" "+repLine.getOrigin())
			continue
		} else if options.OnlyRepeating {
			if !repLine.isUniq() {
				result = append(result, repLine.getOrigin())
			}
			continue
		} else if options.OnlyUnique {
			if repLine.isUniq() {
				result = append(result, repLine.getOrigin())
			}
		}
	}
	return result, nil
}
