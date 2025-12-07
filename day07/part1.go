package day07

import (
	"bufio"
	"log"
	"os"
	// "regexp"
	// "strconv"
	// "strings"
)

func Part1() {
	file, err := os.Open("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	beams := make(map[int]bool)
	totalSplits := 0
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println("current string", currentString)

		// find start
		if len(beams) == 0 {
			for i, char := range currentString {
				if char == 83 { // 83 == S
					beams[i] = true
				}
			}
			// log.Println("beams", beams)
			continue
		}

		newBeams := make(map[int]bool)
		for beamIndex, _ := range beams {
			// log.Println("beamIndex", beamIndex)
			// log.Println("beamIndexalue at string", currentString[beamIndex:beamIndex+1])
			charAtBeamIndex := currentString[beamIndex : beamIndex+1]
			// split the beam
			if charAtBeamIndex == "^" {
				totalSplits += 1
				newBeamIndex := beamIndex - 1
				_, ok := beams[newBeamIndex]
				if !ok {
					newBeams[newBeamIndex] = true
				}
				newBeamIndex = beamIndex + 1
				_, ok = beams[newBeamIndex]
				if !ok {
					newBeams[newBeamIndex] = true
				}
			} else { // or continue beam
				newBeams[beamIndex] = true
			}
		}
		beams = newBeams
		// log.Println("beams", beams)
	}
	// log.Println("total beams", len(beams))
	log.Println("total splits", totalSplits)
}
