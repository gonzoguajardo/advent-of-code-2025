package day06

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numbers := make([][]int, 0)
	operators := make([]string, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println("current string", currentString)
		re := regexp.MustCompile("\\s+")
		if strings.Contains(currentString, "*") {
			split := re.Split(strings.Trim(currentString, " "), -1)
			for _, s := range split {
				operators = append(operators, s)
			}
			break
		}
		split := re.Split(strings.Trim(currentString, " "), -1)
		// log.Println("split", split)
		row := make([]int, 0)
		for _, s := range split {
			n, _ := strconv.Atoi(s)
			row = append(row, n)
		}
		numbers = append(numbers, row)
	}
	// log.Println(numbers)
	// log.Println(operators)

	grandTotal := 0
	for i := 0; i < len(numbers[0]); i++ {
		operator := operators[i]
		isMultiply := false
		total := 0
		if operator == "*" {
			isMultiply = true
			total = 1
		}
		for j := 0; j < len(numbers); j++ {
			if isMultiply {
				total = total * numbers[j][i]
			} else {
				total = total + numbers[j][i]
			}
		}
		// log.Println(total)
		grandTotal += total
	}
	log.Println(grandTotal)
}
