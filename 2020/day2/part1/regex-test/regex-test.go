package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Attempted the same problem but using regexp to split the string for a more cleaner look
// Ended up being in the milliseconds for execution instead of nano
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

		array := regexp.MustCompile("\\s|-|:\\s").Split(s, -1)
		min, err := strconv.Atoi(array[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(array[1])
		if err != nil {
			log.Fatal(err)
		}

		result := strings.Count(array[3], array[2])

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
