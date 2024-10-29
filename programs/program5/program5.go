package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	ManGrep()
}

func ManGrep() {
	after := flag.Int("A", 0, "Kоличество строк, которые будут напечатаны после каждой найденной строки.")
	before := flag.Int("B", 0, "Количество строк, которые будут напечатаны перед каждой найденной строкой.")
	context := flag.Int("C", 0, "Количество строк, которые будут напечатаны вокруг каждой найденной строки.")
	count := flag.Bool("c", false, "Вывод только количества найденных строк.")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр при поиске.")
	invert := flag.Bool("v", false, "Инвертировать результат поиска (выводить строки, которые не соответствуют шаблону).")
	fixed := flag.Bool("F", false, "Интерпретировать шаблон как фиксированную строку (не регулярное выражение).")
	lineNum := flag.Bool("n", false, "Добавлять к каждой строке номер её строки.")

	flag.Parse()

	pattern := flag.Arg(0)

	if pattern == "" {
		fmt.Println("Usage: grep [OPTIONS] PATTERN [FILE]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	matcher := func(line string) bool {
		if *fixed {
			if *ignoreCase {
				return strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			}
			return strings.Contains(line, pattern)
		}
		if *ignoreCase {
			return strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		}
		return strings.Contains(line, pattern)
	}

	printLine := func(line string, lineNumber int) {
		if *lineNum {
			fmt.Printf("%d:%s\n", lineNumber, line)
		} else {
			fmt.Println(line)
		}
	}

	var selectedCount int

	files := flag.Args()[1:]
	if len(files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for lineNumber := 1; scanner.Scan(); lineNumber++ {
			line := scanner.Text()

			if len(lines) > 0 {
				lines = append(lines, line)
			}

			if (matcher(line) && !*invert) || (!matcher(line) && *invert) {
				for _, prevLine := range lines {
					printLine(prevLine, lineNumber-len(lines))
				}

				printLine(line, lineNumber)

				for i := 1; i <= *after; i++ {
					if scanner.Scan() {
						nextLine := scanner.Text()
						printLine(nextLine, lineNumber+i)
					}
				}

				lines = nil

				selectedCount++
			} else if *before > 0 || *context > 0 {
				lines = append(lines, line)
				if len(lines) > *before+*context {
					lines = lines[1:]
				}
			}
		}
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
				continue
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)
			var lines []string
			for lineNumber := 1; scanner.Scan(); lineNumber++ {
				line := scanner.Text()

				if len(lines) > 0 {
					lines = append(lines, line)
				}

				if (matcher(line) && !*invert) || (!matcher(line) && *invert) {
					for _, prevLine := range lines {
						printLine(prevLine, lineNumber-len(lines))
					}

					printLine(line, lineNumber)

					for i := 1; i <= *after; i++ {
						if scanner.Scan() {
							nextLine := scanner.Text()
							printLine(nextLine, lineNumber+i)
						}
					}

					lines = nil

					selectedCount++
				} else if *before > 0 || *context > 0 {
					lines = append(lines, line)
					if len(lines) > *before+*context {
						lines = lines[1:]
					}
				}
			}
		}
	}

	if *count {
		fmt.Printf("Count of selected lines: %d\n", selectedCount)
	}
}
