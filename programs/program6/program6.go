package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ManCut()
}

// ManCut - cut function
func ManCut() {
	// Parsing flags
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	if *fields == "" {
		fmt.Println("Ошибка: не указаны поля для выбора. Используйте флаг -f.")
		return
	}

	// Reading from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, *delimiter)
		if len(parts) == 1 && *delimiter == "\t" {
			parts = strings.Split(line, "   ")
		}

		// Checking for the -s flag condition
		if *separated && len(parts) < 2 {
			continue
		}

		// Printing selected fields to stdout
		outLine := make([]string, 0)
		for _, field := range strings.Split(*fields, ",") {
			index, err := strconv.Atoi(field)
			if err != nil || index < 1 || index > len(parts) {
				continue
			}
			outLine = append(outLine, parts[index-1])
		}
		fmt.Printf("%s\n", strings.Join(outLine, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения STDIN:", err)
		return
	}
}
