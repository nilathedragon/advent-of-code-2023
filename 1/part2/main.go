package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var stringifiedNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	replaceRegex := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)
	splitRegex := regexp.MustCompile(`[^\d]+`)

	content, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, line := range strings.Split(string(content), "\n") {
		// Hack to get around the fact that golang does not do lookahead/lookbehind regex :/
		modifiedLine := replaceRegex.ReplaceAllStringFunc(strings.ToLower(line), func(s string) string {
			split := strings.Split(s, "")
			return split[0] + s + split[len(split)-1]
		})
		modifiedLine = replaceRegex.ReplaceAllStringFunc(modifiedLine, func(s string) string {
			return strconv.Itoa(stringifiedNumbers[s])
		})

		// Trim the line to remove leading and trailing spaces
		modifiedLine = strings.TrimSpace(splitRegex.ReplaceAllString(modifiedLine, " "))
		numbers := strings.Split(modifiedLine, " ")

		if numbers[0] == "" {
			continue
		}

		firstStr, lastStr := numbers[0], numbers[len(numbers)-1]

		lastSplit := strings.Split(lastStr, "")
		number, err := strconv.Atoi(strings.Split(firstStr, "")[0] + lastSplit[len(lastSplit)-1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Line '%s', Modified Line '%s' number: %d\n", line, modifiedLine, number)
		sum += number
	}
	fmt.Println("Sum:", sum)
}
