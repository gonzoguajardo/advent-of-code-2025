package day03

import (
	"bufio"
	"log"
	"os"
	"strconv"
	// "strings"
)

func Part1() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println("current string", currentString)
		max := 0
		firstNumber := 0
		firstNumberIndex := 0
		for i := 0; i < len(currentString); i++ {
			currentDigitString := currentString[i : i+1]
			// log.Println("current digit", currentDigitString)
			currentDigit, _ := strconv.Atoi(currentDigitString)
			// log.Println(currentDigit)
			if currentDigit == 9 {
				max = 9
				firstNumberIndex = i
				break
			}
			if currentDigit > max {
				firstNumberIndex = i
				max = currentDigit
			}
			if i == len(currentString)-2 {
				break
			}
		}
		firstNumber = max
		// log.Println("first digit", firstNumber)
		// log.Println("first digit index", firstNumberIndex)

		secondNumber := 0
		max = 0
		for i := firstNumberIndex + 1; i < len(currentString); i++ {
			currentDigitString := currentString[i : i+1]
			// log.Println("current digit", currentDigitString)
			currentDigit, _ := strconv.Atoi(currentDigitString)
			// log.Println(currentDigit)
			if currentDigit == 9 {
				max = 9
				break
			}
			if currentDigit > max {
				max = currentDigit
			}
		}
		secondNumber = max
		// log.Println("second digit", secondNumber)

		add, _ := strconv.Atoi(strconv.Itoa(firstNumber) + strconv.Itoa(secondNumber))
		total += add
	}
	log.Println("total", total)
}
