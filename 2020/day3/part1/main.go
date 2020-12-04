package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	var count int

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	currPos := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		// Get the charcter of the line we have landed on. Determine if safe or a tree
		if string(s[currPos]) == "#" {
			count++
		}

		// Move over 3
		currPos = currPos + 3
		// Hacky AF but if we go over 31(# of characters in a line). Subtract away 31 to go back to the "beginning"
		if currPos >= 31 {
			currPos = currPos - 31
		}

	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer:", count)

	elapsed := time.Since(start)
	fmt.Printf("Run time %s", elapsed)
}
