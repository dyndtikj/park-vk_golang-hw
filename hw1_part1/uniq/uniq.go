package uniq

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidArgs = fmt.Errorf("invalid arguments")
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
	line       //информация о строке (оригнальная и после применения опций)
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
	//вспомнил дискретку
	return o.IgnoreChars >= 0 && o.IgnoreFields >= 0 &&
		(!o.OnlyUnique && !o.OnlyRepeating ||
			(!o.CountEntries && ((o.OnlyRepeating && !o.OnlyUnique) || (!o.OnlyRepeating && o.OnlyUnique))))
}

// Функция применяет опции к строке, не меняя ее, возвращает копию
func useOptions(line string, options Options) (string, error) {
	result := line
	if options.IgnoreFields > 0 {
		str := strings.Split(result, " ")
		if len(str) < options.IgnoreFields {
			return "", errors.New("Str:\"" + line + "\" dont have " +
				strconv.Itoa(options.IgnoreFields) + " fields")
		}
		result = strings.Join(str[options.IgnoreFields:], " ")
	}
	if options.IgnoreChars > 0 {
		if options.IgnoreChars > len(result) {
			return "", errors.New("Str:" + line + "dont have " +
				strconv.Itoa(options.IgnoreChars) + "letters to ignore")
		}
		if options.IgnoreFields > 0 {
			// При использовании вместе с параметром -f учитываются первые символы после num_fields полей
			// (не учитывая пробел-разделитель после последнего поля).
			result = result[:options.IgnoreChars]
		} else {
			result = result[options.IgnoreChars:]
		}
	}
	if options.IgnoreRegister {
		result = strings.ToLower(result)
	}
	return result, nil
}

func createLines(options Options, input []string) ([]line, error) {
	var result []line
	for _, str := range input {
		modified, err := useOptions(str, options)
		if err != nil {
			return []line{}, fmt.Errorf("failed to use options %w", err)
		}
		result = append(result, line{str, modified})
	}
	return result, nil
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
	var result []string
	if !options.IsValid() {
		return []string{}, ErrInvalidArgs
	}
	lines, err := createLines(options, input)
	if err != nil {
		return []string{}, fmt.Errorf("failed to create lines %w", err)
	}
	repLines := findReplicates(lines)
	for _, repLine := range repLines {
		var outStr string
		if options.CountEntries {
			outStr += strconv.Itoa(int(repLine.count)) + " "
		}
		outStr += repLine.getOrigin()

		if options.OnlyRepeating {
			if !repLine.isUniq() {
				result = append(result, outStr)
			}
			continue
		} else if options.OnlyUnique {
			if repLine.isUniq() {
				result = append(result, outStr)
			}
			continue
		}
		result = append(result, outStr)
	}
	return result, nil
}
