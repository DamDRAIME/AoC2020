// https://adventofcode.com/2020/day/5

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

func binary_find(array string, min, max int, lower, higher rune) int {
	var mid int
	for _, elmt := range array {
		mid = (max-min)/2 + min
		// fmt.Println(string(elmt), " | max: ", max, "; min: ", min, "; mid: ", mid)
		if elmt == lower {
			max = mid
		} else if elmt == higher {
			mid++
			min = mid
		} else {
			log.Fatal("Unexpected element in array: ", string(elmt))
		}
	}
	return mid
}

func get_seat_id(boarding_pass string) int {
	row := binary_find(boarding_pass[:7], 0, 127, 'F', 'B')
	col := binary_find(boarding_pass[7:], 0, 7, 'L', 'R')
	return row*8 + col
}

func part_one() int {
	var highest_seat_id int
	const boarding_passes_path string = "input.txt"
	lines_chnl := make(chan string)
	go Readlines(boarding_passes_path, lines_chnl)
	for boarding_pass := range lines_chnl {
		seat_id := get_seat_id(boarding_pass)
		if seat_id > highest_seat_id {
			highest_seat_id = seat_id
		}
	}
	return highest_seat_id
}

func part_two() int {
	var seat_ids []int
	const boarding_passes_path string = "input.txt"
	lines_chnl := make(chan string)
	go Readlines(boarding_passes_path, lines_chnl)
	for boarding_pass := range lines_chnl {
		seat_id := get_seat_id(boarding_pass)
		seat_ids = append(seat_ids, seat_id)
	}
	sort.Ints(seat_ids)
	for idx, seat_id := range seat_ids {
		if seat_id+1 != seat_ids[idx+1] {
			return seat_id + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(part_one())
	fmt.Println(part_two())
}
