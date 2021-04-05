// https://adventofcode.com/2020/day/6

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func Readlines(path string, lines_chnl chan string) {
	fobj, err := os.Open(path)
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

func is_in(value rune, array []rune) bool {
	return get_index(value, array) >= 0
}

func get_index(value rune, array []rune) int {
	for idx, v := range array {
		if v == value {
			return idx
		}
	}
	return -1
}

func pop(idx int, array []rune) []rune {
	return append(array[:idx], array[idx+1:]...)
}

func pop_indices(indices []int, array []rune) []rune {
	sort.Ints(indices)
	for offset, idx := range indices {
		array = pop(idx-offset, array)
	}
	return array
}

func part_one() int {
	var sum int
	var answers []rune
	const custom_declaration_forms_path string = "input.txt"
	lines_chnl := make(chan string)
	go Readlines(custom_declaration_forms_path, lines_chnl)
	for line := range lines_chnl {
		if line == "" {
			sum += len(answers)
			answers = []rune{}
		} else {
			for _, answer := range line {
				if !is_in(answer, answers) {
					answers = append(answers, answer)
				}
			}
		}
	}
	return sum + len(answers)
}

func part_two() int {
	var sum int
	var answers []rune
	reset := true
	const custom_declaration_forms_path string = "input.txt"
	lines_chnl := make(chan string)
	go Readlines(custom_declaration_forms_path, lines_chnl)
	for line := range lines_chnl {
		if line == "" {
			sum += len(answers)
			answers = []rune{}
			reset = true
		} else if reset {
			for _, answer := range line {
				answers = append(answers, answer)
			}
			reset = false
		} else {
			var indices []int
			for idx, answer := range answers {
				if !is_in(answer, []rune(line)) {
					indices = append(indices, idx)
				}
			}
			answers = pop_indices(indices, answers)
		}
	}
	return sum + len(answers)
}

func main() {
	fmt.Println(part_one())
	fmt.Println(part_two())
}
