package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		splitString := strings.Split(s, "-")
		min, err := strconv.Atoi(splitString[0])
		if err != nil {
			log.Fatal(err)
		}
		splitString = strings.Split(splitString[1], " ")
		max, err := strconv.Atoi(splitString[0])
		if err != nil {
			log.Fatal(err)
		}
		tempLetter := strings.Split(splitString[1], ":")

		result := strings.Count(splitString[2], tempLetter[0])

		if result >= min && result <= max {
			count++
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answer:", count)

	elapsed := time.Since(start)
	fmt.Printf("Run time %s", elapsed)
}
