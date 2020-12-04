package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// Coordinate defines a simple x, y coordinate plane
type Coordinate struct {
	x, y int
}

// CoordinateSystem defines a map of coordinates with its upper x and y bounds
type CoordinateSystem struct {
	coordinates              map[Coordinate]string
	upperboundX, upperboundY int
}

var m map[Coordinate]string

func main() {
	start := time.Now()
	answer := 1

	var slopes []Coordinate
	slopes = append(slopes, Coordinate{x: 1, y: 1})
	slopes = append(slopes, Coordinate{x: 3, y: 1})
	slopes = append(slopes, Coordinate{x: 5, y: 1})
	slopes = append(slopes, Coordinate{x: 7, y: 1})
	slopes = append(slopes, Coordinate{x: 1, y: 2})

	cs := generateCoordinateSystem()

	for _, slope := range slopes {
		count := calculateTreeHits(cs, slope)
		answer = answer * count
	}

	fmt.Println("Answer:", answer)

	elapsed := time.Since(start)
	fmt.Println("Run time:", elapsed)
}

func generateCoordinateSystem() CoordinateSystem {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	currLine := 0
	upperboundX := 0
	upperboundY := 0

	m = make(map[Coordinate]string)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		for currPos := 0; currPos < len(s); currPos++ {
			if upperboundX <= currPos {
				upperboundX = currPos + 1
			}

			tempCoordinate := Coordinate{x: currPos, y: currLine}
			m[tempCoordinate] = string(s[currPos])
		}

		// Increment for next line
		currLine++
		if currLine > upperboundY {
			upperboundY = currLine
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cs := CoordinateSystem{coordinates: m, upperboundX: upperboundX, upperboundY: upperboundY}

	return cs
}

func calculateTreeHits(cs CoordinateSystem, slope Coordinate) int {
	m := cs.coordinates
	upperboundX := cs.upperboundX
	upperboundY := cs.upperboundY
	count := 0
	currY := 0
	currX := 0

	for currY < upperboundY {
		tempCoordinate := Coordinate{x: currX, y: currY}
		// If the current character we are on is a "tree" increment count
		if m[tempCoordinate] == "#" {
			count++
		}

		currX = currX + slope.x
		if currX >= upperboundX {
			currX = currX - upperboundX
		}

		currY = currY + slope.y
	}

	return count
}
