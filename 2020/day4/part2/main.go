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

	return passportList
}

func checkPassports(passportList []passport) int {
	count := 0

	for _, passport := range passportList {
		passportCheck := true

		if passport.byr == 0 || passport.byr < 1920 || passport.byr > 2002 {
			passportCheck = false
		}

		if passport.ecl == "" || (passport.ecl != "amb" && passport.ecl != "blu" && passport.ecl != "brn" && passport.ecl != "gry" && passport.ecl != "grn" && passport.ecl != "hzl" && passport.ecl != "oth") {
			passportCheck = false
		}

		if passport.eyr == 0 || passport.eyr < 2020 || passport.eyr > 2030 {
			passportCheck = false
		}

		res, err := regexp.MatchString("^#[a-f0-9]{6}$", passport.hcl)
		if err != nil {
			log.Fatal(err)
		}
		if passport.hcl == "" || !res || len(passport.hcl) != 7 || passport.hcl[0] != '#' {
			passportCheck = false
		}

		if passport.hgt == "" {
			passportCheck = false
		}
		if strings.Contains(passport.hgt, "cm") {
			tempHgt := strings.Split(passport.hgt, "cm")
			height, err := strconv.Atoi(tempHgt[0])
			if err != nil {
				log.Fatal(err)
			}
			if height < 150 || height > 193 {
				passportCheck = false
			}
		} else if strings.Contains(passport.hgt, "in") {
			tempHgt := strings.Split(passport.hgt, "in")
			height, err := strconv.Atoi(tempHgt[0])
			if err != nil {
				log.Fatal(err)
			}
			if height < 59 || height > 76 {
				passportCheck = false
			}
		} else {
			passportCheck = false
		}

		if passport.iyr == 0 || passport.iyr < 2010 || passport.iyr > 2020 {
			passportCheck = false
		}

		if passport.pid == "" || len(passport.pid) != 9 {
			passportCheck = false
		}

		if passportCheck {
			count++
		}
	}

	return count
}
