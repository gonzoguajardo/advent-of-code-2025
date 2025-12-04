package day04

import (
	"bufio"
	"log"
	"os"
	// "strconv"
	// "strings"
)

func Part2() {
	file, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	M := make([][]string, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println("current string", currentString)
		row := make([]string, 0)
		for i := 0; i < len(currentString); i++ {
			row = append(row, currentString[i:i+1])
		}
		M = append(M, row)
	}
	count := 0
outer:
	for {
		toRemove := make([][]int, 0)
		for rowIndex := 0; rowIndex < len(M); rowIndex++ {
			for colIndex := 0; colIndex < len(M[0]); colIndex++ {

				if M[rowIndex][colIndex] != "@" {
					continue
				}
				if CheckSurround(M, rowIndex, colIndex) < 4 {
					// log.Println("found", rowIndex, colIndex)
					remove := []int{rowIndex, colIndex}
					toRemove = append(toRemove, remove)
					count += 1
				}
			}
		}
		for _, remove := range toRemove {
			M[remove[0]][remove[1]] = "."
		}
		if len(toRemove) == 0 {
			break outer
		}
	}
	log.Println("count", count)
}
