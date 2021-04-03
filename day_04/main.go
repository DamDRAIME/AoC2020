// https://adventofcode.com/2020/day/4

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type year_field struct {
	code     string
	n_digits int
	min      int
	max      int
}

func Readlines(filepath string, lines_chnl chan string) {
	fobj, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fobj)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		lines_chnl <- scanner.Text()
	}
	close(lines_chnl)
}

func has_required_fields(passport map[string]string) bool {
	var required_fields = [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range required_fields {
		_, present := passport[field]
		if !present {
			return false
		}
	}
	return true
}

func is_valid(passport map[string]string) bool {
	// Validating year fields
	byr := year_field{"byr", 4, 1920, 2002}
	iyr := year_field{"iyr", 4, 2010, 2020}
	eyr := year_field{"eyr", 4, 2020, 2030}
	var year_fields = [3]year_field{byr, iyr, eyr}
	for _, field := range year_fields {
		value_str, _ := passport[field.code]
		value, err := strconv.Atoi(value_str)
		if err != nil {
			log.Fatalln(err)
			return false
		}
		if !(value >= field.min && value <= field.max) {
			return false
		}
	}

	// Validating height field
	mins := [2]int{150, 59}
	maxs := [2]int{193, 76}
	units := [2]string{"cm", "in"}
	value_str, _ := passport["hgt"]
	valid := false
	for idx := range units {
		if strings.HasSuffix(value_str, units[idx]) {
			value, err := strconv.Atoi(strings.TrimSuffix(value_str, units[idx]))
			if err != nil {
				log.Fatalln(err)
				return false
			}
			if !(value >= mins[idx] && value <= maxs[idx]) {
				return false
			}
			valid = true
			break
		}
	}
	if !valid {
		return false
	}

	// Validating hair color field
	value_str, _ = passport["hcl"]
	expected_letter_values := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	if strings.HasPrefix(value_str, "#") {
		for _, element := range strings.TrimPrefix(value_str, "#") {
			if unicode.IsDigit(element) {
				continue
			} else if unicode.IsLetter(element) {
				valid := false
				for _, elv := range expected_letter_values {
					if elv == element {
						valid = true
						break
					}
				}
				if !valid {
					return false
				}
			} else {
				return false
			}
		}
	} else {
		return false
	}

	// Validating eye color field
	value_str, _ = passport["ecl"]
	valid_colors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	if !valid_colors[value_str] {
		return false
	}

	// Validating passport id field
	value_str, _ = passport["pid"]
	if len(value_str) != 9 {
		return false
	} else {
		for _, element := range value_str {
			if unicode.IsDigit(element) {
				continue
			} else {
				return false
			}
		}
	}

	return true
}

func part_one() int {
	var count_valid_passports int = 0
	const batch_file_path string = "input.txt"
	lines_chnl := make(chan string)
	go Readlines(batch_file_path, lines_chnl)
	passport := make(map[string]string)
	for line := range lines_chnl {
		if line == "" {
			if has_required_fields(passport) {
				count_valid_passports++
			}
			passport = make(map[string]string)
		} else {
			for _, field_value := range strings.Split(line, " ") {
				field_value := strings.Split(field_value, ":")
				passport[field_value[0]] = field_value[1]
			}
		}
	}
	if has_required_fields(passport) { // To account for the last line
		count_valid_passports++
	}
	return count_valid_passports
}

func part_two() int {
	var count_valid_passports int = 0
	const batch_file_path string = "input.txt"
	lines_chnl := make(chan string)
	go Readlines(batch_file_path, lines_chnl)
	passport := make(map[string]string)
	for line := range lines_chnl {
		if line == "" {
			if has_required_fields(passport) {
				if is_valid(passport) {
					count_valid_passports++
				}
			}
			passport = make(map[string]string)
		} else {
			for _, field_value := range strings.Split(line, " ") {
				field_value := strings.Split(field_value, ":")
				passport[field_value[0]] = field_value[1]
			}
		}
	}
	if has_required_fields(passport) { // To account for the last line
		if is_valid(passport) {
			count_valid_passports++
		}
	}
	return count_valid_passports
}

func main() {
	fmt.Println(part_one())
	fmt.Println(part_two())
}
