package day10

import (
	"bufio"
	"log"
	"os"
	// "regexp"
	// "math"
	// "sort"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lights := make([][]bool, 0)
	// var lightString string
	buttonsForAll := make([][][]int, 0)

	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println(currentString)
		splits := strings.Split(currentString, " ")
		buttonsForLight := make([][]int, 0)
		for _, split := range splits {
			if strings.Contains(split, "[") {
				currentLight := make([]bool, 0)
				lightStringWithoutBrackets := split[1 : len(split)-1]
				for _, lightChar := range lightStringWithoutBrackets {
					if lightChar == 46 { // "."
						currentLight = append(currentLight, false)
					} else {
						currentLight = append(currentLight, true)
					}
				}
				lights = append(lights, currentLight)
			}
			if strings.Contains(split, "(") {
				currentButton := make([]int, 0)
				buttonsWithoutParens := split[1 : len(split)-1]
				buttonStrings := strings.Split(buttonsWithoutParens, ",")
				for _, buttonString := range buttonStrings {
					button, _ := strconv.Atoi(buttonString)
					currentButton = append(currentButton, button)
				}
				buttonsForLight = append(buttonsForLight, currentButton)
			}
		}
		buttonsForAll = append(buttonsForAll, buttonsForLight)
	}
	// log.Println(lights)
	// log.Println(buttons)

	totalK := 0
lightLoop:
	for lightIndex, light := range lights {
		// log.Println("solving for", light)
		currentButtons := buttonsForAll[lightIndex]
		k := 1
		for {
			// log.Println("original light", light)
			combinations := GenerateCombinationsWithRepetition(currentButtons, k)
			// log.Println(combinations)
		combinationLoop:
			for _, combination := range combinations {
				lightCopy := make([]bool, len(light))
				// log.Println("running combination", combination)
				for _, comb := range combination {
					// log.Println("comb", comb)
					for _, toggleLightIndex := range comb {
						// log.Println("toggle light index", toggleLightIndex)
						lightCopy[toggleLightIndex] = !lightCopy[toggleLightIndex]
					}
				}
				// log.Println(lightCopy)
				for checkLightIndex, _ := range lightCopy {
					if lightCopy[checkLightIndex] != light[checkLightIndex] {
						continue combinationLoop
					}
				}
				// log.Println("found valid buttons sequence with k", k)
				totalK += k
				continue lightLoop
			}
			// log.Println("end of k", k)
			k += 1
		}
	}
	log.Println("total K", totalK)
}

// GenerateCombinationsWithRepetition generates all combinations with repetition
// of length k from a set of elements.
func GenerateCombinationsWithRepetition(elements [][]int, k int) [][][]int {
	if k < 0 {
		return [][][]int{}
	}
	if k == 0 {
		return [][][]int{}
	}
	if len(elements) == 0 {
		return [][][]int{}
	}

	var result [][][]int
	var currentCombination [][]int

	var backtrack func(start int, currentCombination [][]int)
	backtrack = func(start int, currentCombination [][]int) {
		if len(currentCombination) == k {
			temp := make([][]int, k)
			copy(temp, currentCombination)
			result = append(result, temp)
			return
		}

		for i := start; i < len(elements); i++ {
			currentCombination = append(currentCombination, elements[i])
			backtrack(i, currentCombination) // 'i' allows repetition
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	backtrack(0, currentCombination)
	return result
}
