package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	re := regexp.MustCompile(`[^\d]+`)

	content, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, line := range strings.Split(string(content), "\n") {
		// Trim the line to remove leading and trailing spaces
		modifiedLine := strings.TrimSpace(re.ReplaceAllString(line, " "))
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

		sum += number
	}
	fmt.Println("Sum:", sum)
}
