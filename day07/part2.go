package day07

import (
	"bufio"
	"log"
	"os"
	// "regexp"
	// "strconv"
	// "strings"
)

func Part2() {
	file, err := os.Open("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	beams := make(map[int]int)
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println("current string", currentString)

		// find start
		if len(beams) == 0 {
			for i, char := range currentString {
				if char == 83 { // 83 == S
					beams[i] = 1
				}
			}
			// log.Println("beams", beams)
			continue
		}

		newBeams := make(map[int]int)
		for beamIndex, _ := range beams {
			charAtBeamIndex := currentString[beamIndex : beamIndex+1]
			currentCount := beams[beamIndex]

			// we want to keep track of how many beams are in a specific index, regardless of whether the beam originiated from a split or a continueation
			// this allows us to "group" the separate dimensions that end up at the same index, without needing to actually do the same work for each one
			// Tricky but the use of maps here really helped as it allowed me to check the index quick and also keep track of how many beams per index

			// split the beam
			if charAtBeamIndex == "^" {
				newBeamIndex := beamIndex - 1
				_, ok := newBeams[newBeamIndex]
				if !ok {
					newBeams[newBeamIndex] = currentCount
				} else {
					newBeams[newBeamIndex] += currentCount
				}
				newBeamIndex = beamIndex + 1
				_, ok = newBeams[newBeamIndex]
				if !ok {
					newBeams[newBeamIndex] = currentCount
				} else {
					newBeams[newBeamIndex] += currentCount
				}
			} else { // or continue beam
				_, ok := newBeams[beamIndex]
				if !ok {
					newBeams[beamIndex] = currentCount
				} else {
					newBeams[beamIndex] += currentCount
				}
			}
		}
		beams = newBeams
		// log.Println("beams", beams)
	}
	// log.Println("beams", beams)

	totalCount := 0
	for _, numberOfBeams := range beams {
		totalCount += numberOfBeams
	}

	log.Println("total count", totalCount)
}
