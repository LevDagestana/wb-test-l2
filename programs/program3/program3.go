package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	SortFile()
}

func SortFile() {
	column := flag.Int("k", 0, "Указание колонки для сортировки")
	numerical := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	if *unique {
		lines = removeDuplicates(lines)
	}

	comparator := createComparator(lines, *column, *numerical)
	sort.SliceStable(lines, func(i, j int) bool {
		if *reverse {
			return !comparator(i, j)
		}
		return comparator(i, j)
	})

	outputFile, err := os.Create("sorted_output.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка записи строки в файл:", err)
			return
		}
	}
}

func createComparator(lines []string, column int, numerical bool) func(int, int) bool {
	return func(i, j int) bool {
		a, b := lines[i], lines[j]

		if column > 0 {
			aCols := strings.Fields(a)
			bCols := strings.Fields(b)

			if column <= len(aCols) && column <= len(bCols) {
				a = aCols[column-1]
				b = bCols[column-1]
			}
		}

		if numerical {
			aNum, errA := strconv.ParseFloat(a, 64)
			bNum, errB := strconv.ParseFloat(b, 64)
			if errA == nil && errB == nil {
				return aNum < bNum
			}
		}

		return a < b
	}
}

func removeDuplicates(lines []string) []string {
	uniqueLines := make(map[string]bool)
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		if !uniqueLines[line] {
			uniqueLines[line] = true
			result = append(result, line)
		}
	}

	return result
}
