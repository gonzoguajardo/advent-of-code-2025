package day03

import (
	"bufio"
	"log"
	"os"
	"strconv"
	// "strings"
)

func Part2() {
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
		startIndex := 0
		number := 0
		currentBatteryString := ""
		for i := 12; i > 0; i-- {
			number, startIndex = findLargest(currentString, startIndex, i)
			// log.Println("number", number)
			// log.Println("index", startIndex)

			currentBatteryString += strconv.Itoa(number)
			startIndex += 1

		}
		// log.Println(currentBatteryString)

		add, _ := strconv.Atoi(currentBatteryString)
		total += add
	}
	log.Println("total", total)
}

func findLargest(currentString string, startIndex int, currentLength int) (int, int) {
	max := 0
	firstNumber := 0
	firstNumberIndex := 0
	for i := startIndex; i < len(currentString); i++ {
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
		if i == len(currentString)-currentLength {
			break
		}
	}
	firstNumber = max
	return firstNumber, firstNumberIndex
}
