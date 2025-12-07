package day06

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	// "strings"
)

func Part2() {
	file, err := os.Open("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// numbers := make([][]int, 0)
	// operators := make([]string, 0)
	strings := make([]string, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println("current string", currentString)
		strings = append(strings, currentString)
	}

	// get lengths?
	operatorString := strings[len(strings)-1]
	numberStrings := strings[0 : len(strings)-1]
	re := regexp.MustCompile(`[*+]`)
	lengths := re.FindAllStringIndex(operatorString, -1)
	// currentIndex := 0
	// log.Println(operatorString)
	grandTotal := 0
	for lenIndex := 0; lenIndex < len(lengths); lenIndex++ {
		operator := operatorString[lengths[lenIndex][0]:lengths[lenIndex][1]]
		total := 0
		isMultiply := false
		if operator == "*" {
			isMultiply = true
			total = 1
		}
		// log.Println(operator)
		start := len(numberStrings[0]) - 1
		if lenIndex+1 < len(lengths) {
			start = lengths[lenIndex+1][0] - 2
		}
		for stringIndex := start; stringIndex >= lengths[lenIndex][0]; stringIndex-- {
			currentNumberString := ""
			// log.Println(stringIndex)
			for _, numberString := range numberStrings {
				stringToAdd := numberString[stringIndex : stringIndex+1]
				if stringToAdd != " " {
					currentNumberString += stringToAdd
				}
			}
			currentNumber, _ := strconv.Atoi(currentNumberString)
			// log.Prinln("curernt number", currentNumber)
			if isMultiply {
				total = total * currentNumber
			} else {
				total += currentNumber
			}
		}
		// log.Println("total", total)
		grandTotal += total
	}
	log.Println(grandTotal)

}
