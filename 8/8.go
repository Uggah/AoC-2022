package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	run_first()
	run_second()
}

func run_first() {
	path := "./input/8.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	x, y := get_matrix_size(path)

	matrix := make([][]int, x)

	for i := 0; i < x; i++ {
		column := make([]int, y)
		matrix[i] = column
	}

	scanner := bufio.NewScanner(file)
	check(err)

	line_count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {

			for i := range line {
				parsed_int, err := strconv.ParseInt(string(line[i]), 10, 64)
				check(err)

				matrix[i][line_count] = int(parsed_int)
			}

			line_count++
		}
	}

	trees_visible := x*y - (x-2)*(y-2)

	for x_1 := 1; x_1 < x-1; x_1++ {
		for y_1 := 1; y_1 < y-1; y_1++ {
			if trees_around_are_lower(x_1, y_1, matrix) {
				trees_visible++
			}
		}
	}

	fmt.Println("First task:", trees_visible)
}

func run_second() {
	path := "./input/8.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	x, y := get_matrix_size(path)

	matrix := make([][]int, x)

	for i := 0; i < x; i++ {
		column := make([]int, y)
		matrix[i] = column
	}

	scanner := bufio.NewScanner(file)
	check(err)

	line_count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {

			for i := range line {
				parsed_int, err := strconv.ParseInt(string(line[i]), 10, 64)
				check(err)

				matrix[i][line_count] = int(parsed_int)
			}

			line_count++
		}
	}

	max := 0

	for x := range matrix {
		for y := range matrix {
			score := get_scenic_score(x, y, matrix)
			if score > max {
				max = score
			}
		}
	}

	fmt.Println("Second task:", max)
}

func get_matrix_size(path string) (int, int) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var x int
	var y int

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			if x == 0 {
				x = len(line)
			}

			y++
		}
	}

	return x, y
}

func trees_around_are_lower(x int, y int, matrix [][]int) bool {
	height := matrix[x][y]

	// Are all trees to the left lower
	left_lower := true
	for x_1 := 0; x_1 < x; x_1++ {
		if matrix[x_1][y] >= height {
			left_lower = false
			break
		}
	}

	// Are all trees to the right lower
	right_lower := true
	for x_1 := x + 1; x_1 < len(matrix); x_1++ {
		if matrix[x_1][y] >= height {

			right_lower = false
			break
		}
	}

	// Are all trees to the top lower
	top_lower := true
	for y_1 := 0; y_1 < y; y_1++ {
		if matrix[x][y_1] >= height {
			top_lower = false
			break
		}
	}

	// Are all trees to the bottom lower
	bottom_lower := true
	for y_1 := y + 1; y_1 < len(matrix[x]); y_1++ {
		if matrix[x][y_1] >= height {
			bottom_lower = false
			break
		}
	}

	return left_lower || right_lower || top_lower || bottom_lower
}

func get_scenic_score(x int, y int, matrix [][]int) int {
	height := matrix[x][y]

	left_trees := 0
	for x_1 := x; x_1 >= 0; x_1-- {
		if x_1 == x {
			continue
		}

		left_trees++
		if matrix[x_1][y] >= height {
			break
		}
	}

	right_trees := 0
	for x_1 := x + 1; x_1 < len(matrix); x_1++ {
		right_trees++
		if matrix[x_1][y] >= height {
			break
		}
	}

	top_trees := 0
	for y_1 := y; y_1 >= 0; y_1-- {
		if y_1 == y {
			continue
		}

		top_trees++
		if matrix[x][y_1] >= height {
			break
		}
	}

	bottom_trees := 0
	for y_1 := y + 1; y_1 < len(matrix[x]); y_1++ {
		bottom_trees++
		if matrix[x][y_1] >= height {
			break
		}
	}

	return left_trees * bottom_trees * right_trees * top_trees
}
