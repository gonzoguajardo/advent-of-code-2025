package day02

import (
	"bufio"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := big.NewInt(0)
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println("current string", currentString)
		split := strings.Split(currentString, ",")
		// log.Println(split)
		for _, set := range split {
			splitSet := strings.Split(set, "-")
			start, _ := strconv.Atoi(splitSet[0])
			end, _ := strconv.Atoi(splitSet[1])
		iLoop:
			for i := start; i <= end; i++ {
				stringI := strconv.Itoa(i)

				for j := 1; j <= len(stringI)/2; j++ {
					currentCheck := stringI[:j]
					// log.Println("building out string", currentCheck)
					checkValid := strings.ReplaceAll(stringI, currentCheck, "")
					if len(checkValid) == 0 {
						// log.Println("found invalid", stringI)
						add := big.NewInt(0)
						add.SetInt64(int64(i)) // big int not really needed
						total = total.Add(total, add)
						continue iLoop
					}
				}
			}
		}
	}
	log.Println("total", total)
}
