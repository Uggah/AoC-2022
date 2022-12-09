package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	debug := false

	path := "./input/9.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	instructions := make([]string, 0)

	var curr_x int
	var curr_y int

	var max_x int
	var max_y int
	var min_x int
	var min_y int

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			instructions = append(instructions, line)

			splitted_line := strings.Split(line, " ")

			if splitted_line[0] == "R" {
				parsed_int, err := strconv.ParseInt(splitted_line[1], 10, 64)
				check(err)

				curr_x += int(parsed_int)
			}

			if splitted_line[0] == "L" {
				parsed_int, err := strconv.ParseInt(splitted_line[1], 10, 64)
				check(err)

				curr_x -= int(parsed_int)
			}

			if splitted_line[0] == "U" {
				parsed_int, err := strconv.ParseInt(splitted_line[1], 10, 64)
				check(err)

				curr_y += int(parsed_int)
			}

			if splitted_line[0] == "D" {
				parsed_int, err := strconv.ParseInt(splitted_line[1], 10, 64)
				check(err)

				curr_y -= int(parsed_int)
			}

			if curr_x > max_x {
				max_x = curr_x
			} else if curr_x < min_x {
				min_x = curr_x
			}

			if curr_y > max_y {
				max_y = curr_y
			} else if curr_y < min_y {
				min_y = curr_y
			}
		}
	}

	x := max_x - min_x
	y := max_y - min_y

	visited_matrix := make([][]bool, x+1)

	for i := range visited_matrix {
		slice := make([]bool, y+1)

		for j := range slice {
			slice[j] = false
		}

		visited_matrix[i] = slice
	}

	var h_x int64
	var t_x int64
	var h_y int64
	var t_y int64

	for i := range instructions {
		instruction := strings.Split(instructions[i], " ")

		direction := instruction[0]
		amount, err := strconv.ParseInt(instruction[1], 10, 64)
		check(err)

		if debug {
			fmt.Println("==", instructions[i], "==")
			fmt.Println()
		}

		for j := 0; j < int(amount); j++ {
			switch direction {
			case "U":
				h_y++
			case "D":
				h_y--
			case "R":
				h_x++
			case "L":
				h_x--
			}

			// .H
			// ..
			// T.
			if h_x == t_x+1 && h_y > t_y+1 {
				t_x++
				t_y = h_y - 1
			} else

			// H.
			// ..
			// .T
			if h_x == t_x-1 && h_y > t_y+1 {
				t_x--
				t_y = h_y - 1
			} else

			// T.
			// ..
			// .H
			if h_x == t_x+1 && h_y < t_y-1 {
				t_x++
				t_y = h_y + 1
			} else

			// .T
			// ..
			// H.
			if h_x == t_x-1 && h_y < t_y-1 {
				t_x--
				t_y = h_y + 1
			} else

			// .H.
			// ...
			// .T.
			if h_x == t_x && h_y > t_y {
				t_y = h_y - 1
			} else

			// .T.
			// ...
			// .H.
			if h_x == t_x && h_y < t_y {
				t_y = h_y + 1
			} else

			// ...
			// T.H
			// ...
			if h_x > t_x && h_y == t_y {
				t_x = h_x - 1
			} else

			// ...
			// H.T
			// ...
			if h_x < t_x && h_y == t_y {
				t_x = h_x + 1
			} else

			// H..
			// ..T
			if h_x < t_x-1 && h_y > t_y {
				t_y++
				t_x = h_x + 1
			} else

			// ..H
			// T..
			if h_x > t_x+1 && h_y > t_y {
				t_y++
				t_x = h_x - 1
			} else

			// T..
			// ..H
			if h_x > t_x+1 && h_y < t_y {
				t_y--
				t_x = h_x - 1
			} else

			// ..T
			// H..
			if h_x < t_x-1 && h_y < t_y {
				t_y--
				t_x = h_x + 1
			}

			t_x_c := int(t_x) + Abs(min_x)
			t_y_c := int(t_y) + Abs(min_y)

			visited_matrix[t_x_c][t_y_c] = true

			if debug {
				h_x_c := int(h_x) + Abs(min_x)
				h_y_c := int(h_y) + Abs(min_y)
				print_matrix(h_x_c, h_y_c, t_x_c, t_y_c, x, y)
			}
		}

	}

	counter := count_true(visited_matrix)

	fmt.Println("First task:", counter)
}

func run_second() {
	debug := false

	path := "./input/9.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	instructions := make([]string, 0)

	var curr_x int
	var curr_y int

	var max_x int
	var max_y int
	var min_x int
	var min_y int

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			instructions = append(instructions, line)

			splitted_line := strings.Split(line, " ")

			if splitted_line[0] == "R" {
				parsed_int, err := strconv.ParseInt(splitted_line[1], 10, 64)
				check(err)

				curr_x += int(parsed_int)
			}

			if splitted_line[0] == "L" {
				parsed_int, err := strconv.ParseInt(splitted_line[1], 10, 64)
				check(err)

				curr_x -= int(parsed_int)
			}

			if splitted_line[0] == "U" {
				parsed_int, err := strconv.ParseInt(splitted_line[1], 10, 64)
				check(err)

				curr_y += int(parsed_int)
			}

			if splitted_line[0] == "D" {
				parsed_int, err := strconv.ParseInt(splitted_line[1], 10, 64)
				check(err)

				curr_y -= int(parsed_int)
			}

			if curr_x > max_x {
				max_x = curr_x
			} else if curr_x < min_x {
				min_x = curr_x
			}

			if curr_y > max_y {
				max_y = curr_y
			} else if curr_y < min_y {
				min_y = curr_y
			}
		}
	}

	x := max_x - min_x
	y := max_y - min_y

	visited_matrix := make([][]bool, x+1)

	for i := range visited_matrix {
		slice := make([]bool, y+1)

		for j := range slice {
			slice[j] = false
		}

		visited_matrix[i] = slice
	}

	knot_map := make(map[int][]int)
	knot_map[0] = make([]int, 2)
	knot_map[1] = make([]int, 2)
	knot_map[2] = make([]int, 2)
	knot_map[3] = make([]int, 2)
	knot_map[4] = make([]int, 2)
	knot_map[5] = make([]int, 2)
	knot_map[6] = make([]int, 2)
	knot_map[7] = make([]int, 2)
	knot_map[8] = make([]int, 2)
	knot_map[9] = make([]int, 2)

	for i := range instructions {
		instruction := strings.Split(instructions[i], " ")

		direction := instruction[0]
		amount, err := strconv.ParseInt(instruction[1], 10, 64)
		check(err)

		if debug {
			fmt.Println("==", instructions[i], "==")
			fmt.Println()
		}

		for j := 0; j < int(amount); j++ {
			for k := 0; k < 10; k++ {

				if k == 0 {
					switch direction {
					case "U":
						knot_map[k][1] = knot_map[k][1] + 1
					case "D":
						knot_map[k][1] = knot_map[k][1] - 1
					case "R":
						knot_map[k][0] = knot_map[k][0] + 1
					case "L":
						knot_map[k][0] = knot_map[k][0] - 1
					}

					continue
				}

				h_x := knot_map[k-1][0]
				h_y := knot_map[k-1][1]
				t_x := knot_map[k][0]
				t_y := knot_map[k][1]

				// .H
				// ..
				// T.
				if h_x == t_x+1 && h_y > t_y+1 {
					t_x++
					t_y = h_y - 1
				} else

				// H.
				// ..
				// .T
				if h_x == t_x-1 && h_y > t_y+1 {
					t_x--
					t_y = h_y - 1
				} else

				// T.
				// ..
				// .H
				if h_x == t_x+1 && h_y < t_y-1 {
					t_x++
					t_y = h_y + 1
				} else

				// .T
				// ..
				// H.
				if h_x == t_x-1 && h_y < t_y-1 {
					t_x--
					t_y = h_y + 1
				} else

				// .H.
				// ...
				// .T.
				if h_x == t_x && h_y > t_y {
					t_y = h_y - 1
				} else

				// .T.
				// ...
				// .H.
				if h_x == t_x && h_y < t_y {
					t_y = h_y + 1
				} else

				// ...
				// T.H
				// ...
				if h_x > t_x && h_y == t_y {
					t_x = h_x - 1
				} else

				// ...
				// H.T
				// ...
				if h_x < t_x && h_y == t_y {
					t_x = h_x + 1
				} else

				// H..
				// ..T
				if h_x < t_x-1 && h_y > t_y {
					t_y++
					t_x = h_x + 1
				} else

				// ..H
				// T..
				if h_x > t_x+1 && h_y > t_y {
					t_y++
					t_x = h_x - 1
				} else

				// T..
				// ..H
				if h_x > t_x+1 && h_y < t_y {
					t_y--
					t_x = h_x - 1
				} else

				// ..T
				// H..
				if h_x < t_x-1 && h_y < t_y {
					t_y--
					t_x = h_x + 1
				}

				knot_map[k][0] = t_x
				knot_map[k][1] = t_y

				if k == 9 {
					t_x_c := int(t_x) + Abs(min_x)
					t_y_c := int(t_y) + Abs(min_y)

					visited_matrix[t_x_c][t_y_c] = true
				}
			}

			if debug {
				print_long_matrix(knot_map, x, y)
			}
		}

	}

	counter := count_true(visited_matrix)

	fmt.Println("Second task:", counter)
}

func count_true(matrix [][]bool) int {
	counter := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] {
				counter++
			}
		}
	}

	return counter
}

func print_matrix(h_x int, h_y int, t_x int, t_y int, x int, y int) {
	matrix := make([][]byte, y+1)

	for i := range matrix {
		slice := make([]byte, x+1)

		for j := range slice {
			slice[j] = '.'
		}

		if i == 0 {
			slice[0] = 's'
		}

		matrix[i] = slice
	}

	matrix[h_y][h_x] = 'H'
	matrix[t_y][t_x] = 'T'

	for i := range matrix {
		for j := range matrix[i] {
			char := string(matrix[y-i][j])
			fmt.Print(char)
		}

		fmt.Print("\n")
	}

	fmt.Print("\n")

}

func print_long_matrix(knot_map map[int][]int, x int, y int) {
	matrix := make([][]byte, y+1)

	for i := range matrix {
		slice := make([]byte, x+1)

		for j := range slice {
			slice[j] = '.'
		}

		if i == 0 {
			slice[0] = 's'
		}

		matrix[i] = slice
	}

	for i := 9; i >= 0; i-- {
		x := knot_map[i][0]
		y := knot_map[i][1]

		if i == 0 {
			matrix[y][x] = 'H'
		} else {
			matrix[y][x] = strconv.Itoa(i)[0]
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			char := string(matrix[y-i][j])
			fmt.Print(char)
		}

		fmt.Print("\n")
	}

	fmt.Print("\n")

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
