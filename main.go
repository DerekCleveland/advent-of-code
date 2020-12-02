package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

// ErrNoAnswer is used when the input is looped through with no answer found
var ErrNoAnswer = errors.New("No answer found")

func main() {
	// start := time.Now()

	var expenseValues []int

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		expenseValues = append(expenseValues, x)
	}

	for _, value := range expenseValues {
		answer, err := checkValue(value, expenseValues)
		if err == nil {
			fmt.Println("Answer:", answer)
			break
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// elapsed := time.Since(start)
	// fmt.Printf("Run time %s", elapsed)
}

// Check value takes in a current value and a array of values to see if any pairs sum to 2020
func checkValue(value int, valueArray []int) (int, error) {
	for _, currValue := range valueArray {
		if currValue+value == 2020 {
			return currValue * value, nil
		}
	}

	return 0, ErrNoAnswer
}
