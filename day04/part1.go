package day04

import (
	"bufio"
	"log"
	"os"
	// "strconv"
	// "strings"
)

func Part1() {
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
	for rowIndex := 0; rowIndex < len(M); rowIndex++ {
		for colIndex := 0; colIndex < len(M[0]); colIndex++ {

			if M[rowIndex][colIndex] != "@" {
				continue
			}
			if checkSurround(M, rowIndex, colIndex) < 4 {
				// log.Println("found", rowIndex, colIndex)
				count += 1
			}
		}
	}
	log.Println("count", count)
}

func checkSurround(M [][]string, rowIndex int, colIndex int) int {
	if M[rowIndex][colIndex] != "@" {
		return 0
	}
	found := 0
	for surround := 0; surround < 8; surround++ {
		checkRowIndex := rowIndex
		checkColIndex := colIndex
		switch surround {
		case 0:
			checkRowIndex = rowIndex - 1
			checkColIndex = colIndex - 1
			if validRowIndex(M, checkRowIndex) && validColIndex(M, checkColIndex) && M[checkRowIndex][checkColIndex] == "@" {
				// log.Println("found at 0")
				found += 1
			}
		case 1:
			checkRowIndex = rowIndex - 1
			if validRowIndex(M, checkRowIndex) && validColIndex(M, checkColIndex) && M[checkRowIndex][checkColIndex] == "@" {
				// log.Println("found at 1")
				found += 1
			}
		case 2:
			checkRowIndex = rowIndex - 1
			checkColIndex = colIndex + 1
			if validRowIndex(M, checkRowIndex) && validColIndex(M, checkColIndex) && M[checkRowIndex][checkColIndex] == "@" {
				// log.Println("found at 2")
				found += 1
			}
		case 3:
			checkColIndex = colIndex - 1
			if validRowIndex(M, checkRowIndex) && validColIndex(M, checkColIndex) && M[checkRowIndex][checkColIndex] == "@" {
				// log.Println("found at 3")
				found += 1
			}
		case 4:
			checkColIndex = colIndex + 1
			if validRowIndex(M, checkRowIndex) && validColIndex(M, checkColIndex) && M[checkRowIndex][checkColIndex] == "@" {
				// log.Println("found at 4")
				found += 1
			}
		case 5:
			checkRowIndex = rowIndex + 1
			checkColIndex = colIndex - 1
			if validRowIndex(M, checkRowIndex) && validColIndex(M, checkColIndex) && M[checkRowIndex][checkColIndex] == "@" {
				// log.Println("found at 5")
				found += 1
			}
		case 6:
			checkRowIndex = rowIndex + 1
			if validRowIndex(M, checkRowIndex) && validColIndex(M, checkColIndex) && M[checkRowIndex][checkColIndex] == "@" {
				// log.Println("found at 6")
				found += 1
			}
		case 7:
			checkRowIndex = rowIndex + 1
			checkColIndex = colIndex + 1
			if validRowIndex(M, checkRowIndex) && validColIndex(M, checkColIndex) && M[checkRowIndex][checkColIndex] == "@" {
				// log.Println("found at 7")
				found += 1
			}
		}
	}
	return found
}

func validRowIndex(M [][]string, rowIndex int) bool {
	if rowIndex < 0 {
		return false
	}
	if rowIndex >= len(M) {
		return false
	}
	return true
}

func validColIndex(M [][]string, colIndex int) bool {
	if colIndex < 0 {
		return false
	}
	if colIndex >= len(M[0]) {
		return false
	}
	return true
}
