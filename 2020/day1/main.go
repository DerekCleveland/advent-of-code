package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// ErrNoAnswer is used when the input is looped through with no answer found
var ErrNoAnswer = errors.New("No answer found")

func main() {
	start := time.Now()

	var values []int
	sum := 2020

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, i)
	}

	answer, err := lookForPairSum(sum, values)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Answer 1:", answer)

	answer, err = lookForTripletSum(sum, values)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Answer 2:", answer)

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Run time %s", elapsed)
}

// Checks to see if a pair sums to 2020 and returns the answer
func lookForPairSum(sum int, entries []int) (int, error) {
	for i := range entries {
		for j := i; j < len(entries); j++ {
			if entries[i]+entries[j] == sum {
				return entries[i] * entries[j], nil
			}
		}
	}

	return 0, ErrNoAnswer
}

// Check to see if a triplet sums to 2020 and returns the answer
func lookForTripletSum(sum int, entries []int) (int, error) {
	for _, i := range entries {
		for _, j := range entries {
			for _, k := range entries {
				if i+j+k == sum {
					return i * j * k, nil
				}
			}
		}
	}

	return 0, ErrNoAnswer
}
