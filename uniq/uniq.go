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
func useOptions(line string, options Options) (result string, err error) {
	result = line
	if options.IgnoreFields > 0 {
		str := strings.Split(result, " ")
		if len(str) < options.IgnoreFields {
			return result, errors.New("Str:\"" + line + "\" dont have " +
				strconv.Itoa(options.IgnoreFields) + " fields")
		}
		result = strings.Join(str[options.IgnoreFields:], " ")
	}
	if options.IgnoreChars > 0 {
		if options.IgnoreFields > 0 {
			// При использовании вместе с параметром -f учитываются первые символы после num_fields полей
			// (не учитывая пробел-разделитель после последнего поля).
			if options.IgnoreChars > len(result) {
				return result, errors.New("Str:\"" + line + "\" dont have " +
					strconv.Itoa(options.IgnoreChars) + " letters after applying ignore " +
					strconv.Itoa(options.IgnoreFields) + " fields option")
			}
			result = result[:options.IgnoreChars]
		} else {
			result = result[options.IgnoreChars:]
		}
	}
	if options.IgnoreRegister {
		result = strings.ToLower(result)
	}
	return
}

func createLines(options Options, input []string) (result []line, err error) {
	for _, str := range input {
		var modified string
		modified, err = useOptions(str, options)
		if err != nil {
			return
		}
		result = append(result, line{str, modified})
	}
	return
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

func Uniq(options Options, input []string) (result []string, err error) {
	if !options.IsValid() {
		return result, errors.New("invalid arguments")
	}
	lines, err := createLines(options, input)
	if err != nil {
		return
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
