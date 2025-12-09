package day09

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
	file, err := os.Open("day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	points := make([][]int, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println(currentString)

		split := strings.Split(currentString, ",")
		currentPoint := make([]int, 0)
		for _, numString := range split {
			num, _ := strconv.Atoi(numString)
			currentPoint = append(currentPoint, num)
		}
		points = append(points, currentPoint)
	}
	// log.Println(points)

	// find distances between all points and create map
	areaMap := make(map[[2]int]float64)
	for i := 0; i < len(points); i++ {
		firstPoint := points[i]
		for j := i + 1; j < len(points); j++ {
			secondPoint := points[j]
			xDiff := math.Max(1, math.Abs(float64(firstPoint[0]-secondPoint[0]))) + 1
			yDiff := math.Max(1, math.Abs(float64(firstPoint[1]-secondPoint[1]))) + 1
			// distance := math.Sqrt(math.Pow(float64(firstPoint[0]-secondPoint[0]), 2) + math.Pow(float64(firstPoint[1]-secondPoint[1]), 2) + math.Pow(float64(firstPoint[2]-secondPoint[2]), 2))
			areaMap[[2]int{i, j}] = xDiff * yDiff
		}
	}
	// log.Println(areaMap)

	// sort by area between points
	var sortedPairs []KeyValuePair
	for k, v := range areaMap { // myMap is your original map
		sortedPairs = append(sortedPairs, KeyValuePair{Key: k, Value: v})
	}
	sort.Slice(sortedPairs, func(i, j int) bool {
		return sortedPairs[i].Value > sortedPairs[j].Value
	})
	log.Println(sortedPairs[0].Value)
}

type KeyValuePair struct {
	Key   [2]int
	Value float64
}
