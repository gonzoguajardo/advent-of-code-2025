package day08

import (
	"bufio"
	"log"
	"os"
	// "regexp"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	points := make([][]int, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		split := strings.Split(currentString, ",")
		currentPoint := make([]int, 0)
		for _, numString := range split {
			num, _ := strconv.Atoi(numString)
			currentPoint = append(currentPoint, num)
		}
		points = append(points, currentPoint)
	}

	// find distances between all points and create map
	distanceMap := make(map[[2]int]float64)
	for i := 0; i < len(points); i++ {
		firstPoint := points[i]
		for j := i + 1; j < len(points); j++ {
			secondPoint := points[j]
			distance := math.Sqrt(math.Pow(float64(firstPoint[0]-secondPoint[0]), 2) + math.Pow(float64(firstPoint[1]-secondPoint[1]), 2) + math.Pow(float64(firstPoint[2]-secondPoint[2]), 2))
			distanceMap[[2]int{i, j}] = distance
		}
	}

	// sort by distance between points
	var sortedPairs []KeyValuePair
	for k, v := range distanceMap { // myMap is your original map
		sortedPairs = append(sortedPairs, KeyValuePair{Key: k, Value: v})
	}
	sort.Slice(sortedPairs, func(i, j int) bool {
		return sortedPairs[i].Value < sortedPairs[j].Value
	})
	// log.Println(sortedPairs)

	// add to circuits
	circuits := make([]map[int]bool, 0)
sortedPairsLoop:
	for i, pair := range sortedPairs {
		// log.Println(pair)
		currentPoints := pair.Key

		foundCircuits := make([]int, 0)
		for circuitIndex, circuit := range circuits {
			_, ok0 := circuit[currentPoints[0]]
			_, ok1 := circuit[currentPoints[1]]
			if ok0 || ok1 {
				foundCircuits = append(foundCircuits, circuitIndex)
			}
		}

		if len(foundCircuits) == 0 {
			newCircuit := make(map[int]bool)
			newCircuit[currentPoints[0]] = true
			newCircuit[currentPoints[1]] = true
			circuits = append(circuits, newCircuit)
		} else if len(foundCircuits) == 1 {
			circuits[foundCircuits[0]][currentPoints[0]] = true
			circuits[foundCircuits[0]][currentPoints[1]] = true
		} else if len(foundCircuits) == 2 { // combine circuits if needed
			// log.Println("found multiple circuits")
			combinedCircuit := make(map[int]bool)
			for k, v := range circuits[foundCircuits[0]] {
				combinedCircuit[k] = v
			}
			for k, v := range circuits[foundCircuits[1]] {
				combinedCircuit[k] = v
			}
			// log.Println("combined circuit", combinedCircuit)
			circuits = append(circuits, combinedCircuit)
			// delete existing circuits since it is now merged
			circuits = append(circuits[:foundCircuits[0]], circuits[foundCircuits[0]+1:]...)
			circuits = append(circuits[:foundCircuits[1]-1], circuits[foundCircuits[1]:]...)
		} else {
			log.Println("this should not happen")
			break sortedPairsLoop

		}

		if i == 999 {
			log.Println("reached i==9")
			break
		}

	}
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	log.Println(len(circuits[0]) * len(circuits[1]) * len(circuits[2]))
}

type KeyValuePair struct {
	Key   [2]int
	Value float64
}
