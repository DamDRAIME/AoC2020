// https://adventofcode.com/2020/day/1

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func part_one() int {
	const expense_report_fp = "input.txt"
	lines_chnl := make(chan string)
	go Readlines(expense_report_fp, lines_chnl)

	const total = 2020
	complements := make(map[int]bool)

	for line := range lines_chnl {
		value, _ := strconv.Atoi(line)
		if complements[value] {
			return value * (total - value)
		} else {
			complements[(total - value)] = true
		}
	}
	return -1
}

func part_two() int {
	const expense_report_fp = "input.txt"
	lines_chnl := make(chan string)
	go Readlines(expense_report_fp, lines_chnl)

	const total = 2020
	var values []int

	i := 0
	for line := range lines_chnl {
		value, _ := strconv.Atoi(line)
		values = append(values, value)
		i++
	}
	n_values := i

	for idx1, value1 := range values[0 : n_values-2] {
		for idx2, value2 := range values[idx1 : n_values-1] {
			value12 := value1 + value2
			for _, value3 := range values[idx2:] {
				if value12+value3 == total {
					return value1 * value2 * value3
				}
			}
		}
	}
	return -1
}

func main() {
	// fmt.Println(part_one())
	fmt.Println(part_two())
}
