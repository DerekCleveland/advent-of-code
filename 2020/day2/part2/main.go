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
		pos1, err := strconv.Atoi(splitString[0])
		if err != nil {
			log.Fatal(err)
		}
		splitString = strings.Split(splitString[1], " ")
		pos2, err := strconv.Atoi(splitString[0])
		if err != nil {
			log.Fatal(err)
		}
		tempLetter := strings.Split(splitString[1], ":")

		letter := tempLetter[0]
		password := splitString[2]

		if (string(password[pos1-1]) == letter) != (string(password[pos2-1]) == letter) {
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
