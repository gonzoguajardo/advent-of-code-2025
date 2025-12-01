package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Part1() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentPointer := 50
	numberOfZeros := 0
	for scanner.Scan() {
		currentString := scanner.Text()
		direction := currentString[0:1]
		distance, _ := strconv.Atoi(currentString[1:])
		distance = distance % 100
		// log.Println("current string", currentString)
		// log.Println("direction", direction)
		// log.Println("distance", distance)
		if direction == "L" {
			currentPointer = currentPointer - distance
			if currentPointer < 0 {
				currentPointer = 100 + currentPointer
			}
		}
		if direction == "R" {
			currentPointer = currentPointer + distance
			if currentPointer > 99 {
				currentPointer = currentPointer - 100
			}
		}
		if currentPointer == 0 {
			numberOfZeros += 1
		}
		// log.Println("current pointer", currentPointer)
	}
	log.Println("number of zeros", numberOfZeros)
}
