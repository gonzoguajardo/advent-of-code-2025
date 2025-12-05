package day05

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ranges := make([][]int, 0)
	parsingRanges := true
	freshCount := 0
scanner:
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println("current string", currentString)
		if currentString == "" {
			parsingRanges = false
			continue
		}
		if parsingRanges {
			split := strings.Split(currentString, "-")
			start, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			ranges = append(ranges, []int{start, end})
		} else {
			check, _ := strconv.Atoi(currentString)
			for _, r := range ranges {
				if check >= r[0] && check <= r[1] {
					freshCount += 1
					continue scanner
				}
			}
		}
	}
	log.Println(freshCount)
}
