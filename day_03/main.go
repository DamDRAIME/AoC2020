// https://adventofcode.com/2020/day/3

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

func parse_char(char rune) int {
	var tree rune = []rune("#")[0]
	if char == tree {
		return 1
	} else {
		return 0
	}
}

func get_geology_map(geology_map_path string) [][]int {
	var geology_map [][]int
	lines_chnl := make(chan string)
	go Readlines(geology_map_path, lines_chnl)
	for line := range lines_chnl {
		var line_bit []int
		for _, char := range line {
			line_bit = append(line_bit, parse_char(char))
		}
		geology_map = append(geology_map, line_bit)
	}
	return geology_map

}

func min(array [5]int) int {
	min := array[0]
	for _, v := range array {
		if v < min {
			min = v
		}
	}
	return min
}

func is_div_by(num, den int) bool {
	remainder := int(math.Mod(float64(num), float64(den)))
	if remainder == 0 {
		return true
	} else {
		return false
	}
}

func mult(array [5]int) int {
	value := 1
	for _, v := range array {
		value *= v
	}
	return value
}

type rational_slope struct {
	x int
	y int
}

func part_one() int {
	var count_trees int = 0
	const geology_map_path = "input.txt"
	geology_map := get_geology_map(geology_map_path)
	n_cols := len(geology_map[0])
	n_rows := len(geology_map)
	slope := rational_slope{3, 1}
	pos_x := 0
	for pos_y := slope.y; pos_y < n_rows; pos_y += slope.y {
		pos_x += slope.x
		if pos_x >= n_cols {
			pos_x = pos_x - n_cols
		}
		count_trees += geology_map[pos_y][pos_x]
	}
	return count_trees
}

func part_two() int {
	var count_trees [5]int
	const geology_map_path = "input.txt"
	geology_map := get_geology_map(geology_map_path)
	n_cols := len(geology_map[0])
	n_rows := len(geology_map)

	var slopes [5]rational_slope
	slopes_x := [5]int{1, 3, 5, 7, 1}
	slopes_y := [5]int{1, 1, 1, 1, 2}
	for idx := range slopes_x {
		slopes[idx] = rational_slope{slopes_x[idx], slopes_y[idx]}
	}
	min_increment_y := min(slopes_y)

	var pos_x [5]int
	for pos_y := min_increment_y; pos_y < n_rows; pos_y += min_increment_y {
		for idx, slope := range slopes {
			if is_div_by(pos_y, slope.y) {
				pos_x_ := pos_x[idx]
				pos_x_ += slope.x
				if pos_x_ >= n_cols {
					pos_x_ = pos_x_ - n_cols
				}
				count_trees[idx] += geology_map[pos_y][pos_x_]
				pos_x[idx] = pos_x_
			}
		}
	}
	return mult(count_trees)
}

func main() {
	fmt.Println(part_one())
	fmt.Println(part_two())
}
