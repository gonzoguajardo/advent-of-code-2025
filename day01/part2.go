package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Part2() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentPointer := 50
	numberOfZeros := 0
	log.Println("current pointer", currentPointer)
	for scanner.Scan() {
		currentString := scanner.Text()
		direction := currentString[0:1]
		distance, _ := strconv.Atoi(currentString[1:])
		// log.Println("current string", currentString)
		numberOfFullRotations := distance / 100
		if numberOfFullRotations > 0 {
			numberOfZeros += numberOfFullRotations
			// log.Println("number of full rotations", numberOfFullRotations)
		}
		distance = distance % 100
		// log.Println("direction", direction)
		// log.Println("distance", distance)
		newPointer := 0
		if direction == "L" {
			newPointer = currentPointer - distance
			if newPointer < 0 {
				newPointer = 100 + newPointer
				if currentPointer != 0 && newPointer != 0 {
					numberOfZeros += 1
				}
			}
		}
		if direction == "R" {
			newPointer = currentPointer + distance
			if newPointer > 99 {
				newPointer = newPointer - 100
				if currentPointer != 0 && newPointer != 0 {
					numberOfZeros += 1
				}
			}
		}
		currentPointer = newPointer
		if currentPointer == 0 {
			numberOfZeros += 1
		}
		// log.Println("current pointer", currentPointer)
		// log.Println("number of full rotations", numberOfFullRotations)
	}
	log.Println("number of zeros", numberOfZeros)
}
