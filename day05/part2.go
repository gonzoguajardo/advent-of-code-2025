package day05

import (
	"bufio"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part2() {
	file, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ranges := make([][]int, 0)
	parsingRanges := true
	freshCount := big.NewInt(0)
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println("current string", currentString)
		if currentString == "" {
			parsingRanges = false
			break
		}
		if parsingRanges {
			split := strings.Split(currentString, "-")
			start, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			ranges = append(ranges, []int{start, end})
		}
	}

	// sort by start of range
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	// log.Println(ranges)

	// consolidate ranges
	consolidatedRanges := make([][]int, 0)
	consolidatedRanges = append(consolidatedRanges, ranges[0])

	for i := 1; i < len(ranges); i++ {
		lastRange := consolidatedRanges[len(consolidatedRanges)-1]
		if ranges[i][0] >= lastRange[0] && ranges[i][0] <= lastRange[1] {
			if ranges[i][1] > lastRange[1] { // tricky tricky.. you want to make sure you take the max
				lastRange[1] = ranges[i][1]
			}
			consolidatedRanges[len(consolidatedRanges)-1] = lastRange
		} else {
			consolidatedRanges = append(consolidatedRanges, ranges[i])
		}
	}

	// log.Println(consolidatedRanges)
	for _, consolidatedRange := range consolidatedRanges {
		amountToAdd := consolidatedRange[1] - consolidatedRange[0] + 1
		add := big.NewInt(0)
		add.SetInt64(int64(amountToAdd)) // big int not really needed
		freshCount = freshCount.Add(freshCount, add)
	}
	log.Println(freshCount)

}
