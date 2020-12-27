package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()

	switch *methodP {
	case "p1":
		PartOne()
		break
	case "p2":
		PartTwo()
		break
	case "test":
		break
	}
}

func PartOne() {
	input := readInput()

	passports := makePassports(input)

	validPassports := 0

	for _, v := range passports {
		if CheckValidPassportIgnoreCID(v) {
			validPassports++
		}
	}

	fmt.Println("The number of valid passports (ignoring the CID) is", validPassports)
}

func CheckValidPassportIgnoreCID(p Passport) bool {
	if p.BirthYear == "" || p.ExpirationYear == "" || p.EyeColour == "" || p.HairColour == "" || p.Height == "" || p.IssueYear == "" || p.PassportID == "" {
		return false
	}
	return true
}

func PartTwo() {
	input := readInput()

	passports := makePassports(input)

	validPassports := 0

	for _, v := range passports {
		if CheckValidPassportIgnoreCIDPlusStrictChecks(v) {
			validPassports++
		}
	}

	fmt.Println("The number of valid passports (ignoring the CID and with strict checks) is", validPassports)
}

func CheckValidPassportIgnoreCIDPlusStrictChecks(p Passport) bool {

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	birthYear, err := strconv.Atoi(p.BirthYear)
	if err != nil {
		return false
	} else {
		if birthYear < 1920 || birthYear > 2002 {
			return false
		}
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	issueYear, err := strconv.Atoi(p.IssueYear)
	if err != nil {
		return false
	} else {
		if issueYear < 2010 || issueYear > 2020 {
			return false
		}
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	expirationYear, err := strconv.Atoi(p.ExpirationYear)
	if err != nil {
		return false
	} else {
		if expirationYear < 2020 || expirationYear > 2030 {
			return false
		}
	}

	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	regexheight := regexp.MustCompile(`(\d+)(cm|in)`)
	match := regexheight.FindAllStringSubmatch(p.Height, -1)

	if match != nil {
		heightint, err := strconv.Atoi(match[0][1])

		if err != nil {
			return false
		}

		if match[0][2] == "cm" {
			if heightint < 150 || heightint > 193 {
				return false
			}
		} else if match[0][2] == "in" {
			if heightint < 59 || heightint > 76 {
				return false
			}
		}

	} else {
		return false
	}
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hairColourMatch, _ := regexp.MatchString(`#[0-9a-f]{6}$`, p.HairColour)
	if !hairColourMatch {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	eyeColourMatch, _ := regexp.MatchString(`amb|blu|brn|gry|grn|hzl|oth`, p.EyeColour)
	if !eyeColourMatch {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	passportID, _ := regexp.MatchString(`[0-9]{9}$`, p.PassportID)
	if !passportID {
		return false
	}

	// cid (Country ID) - ignored, missing or not.

	return true
}

func makePassports(input []string) []Passport {

	var passports []Passport

	currentBYR := ""
	currentIYR := ""
	currentEYR := ""
	currentHGT := ""
	currentHCL := ""
	currentECL := ""
	currentPID := ""
	currentCID := ""

	for _, v := range input {

		// Finish and reset
		if v == "" {
			passport := Passport{currentBYR, currentIYR, currentEYR, currentHGT, currentHCL, currentECL, currentPID, currentCID}
			passports = append(passports, passport)

			currentBYR = ""
			currentIYR = ""
			currentEYR = ""
			currentHGT = ""
			currentHCL = ""
			currentECL = ""
			currentPID = ""
			currentCID = ""
			continue
		}

		// Gather data
		split := strings.Split(v, " ")
		for _, v := range split {
			splitInner := strings.Split(v, ":")
			switch splitInner[0] {
			case "byr":
				currentBYR = splitInner[1]
				break
			case "iyr":
				currentIYR = splitInner[1]
				break
			case "eyr":
				currentEYR = splitInner[1]
				break
			case "hgt":
				currentHGT = splitInner[1]
				break
			case "hcl":
				currentHCL = splitInner[1]
				break
			case "ecl":
				currentECL = splitInner[1]
				break
			case "pid":
				currentPID = splitInner[1]
				break
			case "cid":
				currentCID = splitInner[1]
				break
			}
		}
	}

	return passports
}

// Passport represents a passport, note that a passport might not have all these fields if invalid
type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColour     string
	EyeColour      string
	PassportID     string
	CountryID      string
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput() []string {

	var input []string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
