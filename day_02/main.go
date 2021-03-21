// https://adventofcode.com/2020/day/2

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func parseline(line string) (int, int, rune, string) {
	line_splitted := strings.Split(line, " ")
	policy_min_max := strings.Split(line_splitted[0], "-")
	policy_min, _ := strconv.Atoi(policy_min_max[0])
	policy_max, _ := strconv.Atoi(policy_min_max[1])
	policy_letter := rune(line_splitted[1][0])
	pwd := line_splitted[2]
	return policy_min, policy_max, policy_letter, pwd
}

func validate_pwd_first_policy(policy_min int, policy_max int, policy_letter rune, pwd string) bool {
	count_letter := 0
	for _, letter := range pwd {
		if letter == policy_letter {
			count_letter++
		}
	}
	if policy_min <= count_letter && count_letter <= policy_max {
		return true
	} else {
		return false
	}
}

func xor(bool1 bool, bool2 bool) bool {
	return bool1 != bool2
}

func validate_pwd_second_policy(idx1 int, idx2 int, policy_letter rune, pwd string) bool {
	idx1--
	idx2--
	if idx1 < 0 || idx2 > len(pwd)-1 {
		return false
	}
	if xor(rune(pwd[idx1]) == policy_letter, rune(pwd[idx2]) == policy_letter) {
		return true
	} else {
		return false
	}
}

func part_one() int {
	const passwords_db = "input.txt"
	lines_chnl := make(chan string)
	go Readlines(passwords_db, lines_chnl)

	var n_valid_pwds int
	for line := range lines_chnl {
		policy_min, policy_max, policy_letter, pwd := parseline(line)
		if validate_pwd_first_policy(policy_min, policy_max, policy_letter, pwd) {
			n_valid_pwds++
		}
	}
	return n_valid_pwds
}

func part_two() int {
	const passwords_db = "input.txt"
	lines_chnl := make(chan string)
	go Readlines(passwords_db, lines_chnl)

	var n_valid_pwds int
	for line := range lines_chnl {
		idx1, idx2, policy_letter, pwd := parseline(line)
		if validate_pwd_second_policy(idx1, idx2, policy_letter, pwd) {
			n_valid_pwds++
		}
	}
	return n_valid_pwds
}

func main() {
	fmt.Println(part_one())
	fmt.Println(part_two())
}
