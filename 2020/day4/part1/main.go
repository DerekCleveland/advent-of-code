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

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)
type passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string // Technically a hex value
	ecl string
	pid string // Cant be a large value or hex value
	cid int
}

func main() {
	start := time.Now()

	passports := readInput()
	answer := checkPassports(passports)

	fmt.Println("Answer:", answer)

	elapsed := time.Since(start)
	fmt.Println("Run time:", elapsed)
}

func readInput() []passport {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var passportList []passport
	var tempPassport passport

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		// For the current line split on spaces to get individual parts of string
		splitString := strings.Split(s, " ")

		for _, part := range splitString {
			if strings.Contains(part, "byr") {
				byr := strings.Split(part, ":")
				tempPassport.byr, err = strconv.Atoi(byr[1])
				if err != nil {
					log.Fatal(err)
				}
			} else if strings.Contains(part, "iyr") {
				iyr := strings.Split(part, ":")
				tempPassport.iyr, err = strconv.Atoi(iyr[1])
				if err != nil {
					log.Fatal(err)
				}
			} else if strings.Contains(part, "eyr") {
				eyr := strings.Split(part, ":")
				tempPassport.eyr, err = strconv.Atoi(eyr[1])
				if err != nil {
					log.Fatal(err)
				}
			} else if strings.Contains(part, "hgt") {
				hgt := strings.Split(part, ":")
				tempPassport.hgt = hgt[1]
			} else if strings.Contains(part, "hcl") {
				hcl := strings.Split(part, ":")
				tempPassport.hcl = hcl[1]
			} else if strings.Contains(part, "ecl") {
				ecl := strings.Split(part, ":")
				tempPassport.ecl = ecl[1]
			} else if strings.Contains(part, "pid") {
				pid := strings.Split(part, ":")
				tempPassport.pid = pid[1]
			} else if strings.Contains(part, "cid") {
				cid := strings.Split(part, ":")
				tempPassport.cid, err = strconv.Atoi(cid[1])
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		// Might have to add an EOF check as well
		if len(s) == 0 {
			// Add current passport and create a new one
			passportList = append(passportList, tempPassport)

			// Clear tempPassport
			tempPassport = passport{}
		}
	}

	// Catch the last passport after we've reached EOF
	passportList = append(passportList, tempPassport)

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("passport count:", len(passportList))
	return passportList
}

func checkPassports(passportList []passport) int {
	count := 0

	for _, passport := range passportList {
		passportCheck := true
		if passport.byr == 0 {
			passportCheck = false
		}

		if passport.ecl == "" {
			passportCheck = false
		}

		if passport.eyr == 0 {
			passportCheck = false
		}

		if passport.hcl == "" {
			passportCheck = false
		}

		if passport.hgt == "" {
			passportCheck = false
		}

		if passport.iyr == 0 {
			passportCheck = false
		}

		if passport.pid == "" {
			passportCheck = false
		}

		if passportCheck {
			count++
		}
	}

	return count
}
