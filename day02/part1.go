package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println("current string", currentString)
		split := strings.Split(currentString, ",")
		// log.Println(split)
		for _, set := range split {
			splitSet := strings.Split(set, "-")
			start, _ := strconv.Atoi(splitSet[0])
			end, _ := strconv.Atoi(splitSet[1])
			for i := start; i <= end; i++ {
				stringI := strconv.Itoa(i)
				if len(stringI)%2 != 0 {
					continue
				}
				firstHalf := stringI[0 : len(stringI)/2]
				secondHalf := stringI[len(stringI)/2:]
				if firstHalf == secondHalf {
					// log.Println("found invalid", stringI)
					total += i
				}
			}
		}
	}
	log.Println("total", total)
}
