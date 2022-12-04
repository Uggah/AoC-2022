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
	file, err := os.Open("./input/4.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	counter := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitted_line := strings.Split(line, ",")
		left_range := make([]int64, 0)
		right_range := make([]int64, 0)

		for i := 0; i < 2; i++ {
			rangeString := splitted_line[i]

			begin, err := strconv.ParseInt(strings.Split(rangeString, "-")[0], 10, 16)
			check(err)
			end, err := strconv.ParseInt(strings.Split(rangeString, "-")[1], 10, 16)
			check(err)

			for j := begin; j <= end; j++ {
				if i == 0 {
					left_range = append(left_range, j)
				} else {
					right_range = append(right_range, j)
				}
			}

		}

		right_contains_left := true
		left_contains_right := true

		// Does right contain all of left?
		for i := range left_range {
			if !does_contain(right_range, left_range[i]) {
				right_contains_left = false
				break
			}
		}

		// Does left contain all of right?
		for i := range right_range {
			if !does_contain(left_range, right_range[i]) {
				left_contains_right = false
				break
			}
		}

		if left_contains_right || right_contains_left {
			counter++
		}

	}

	fmt.Println("First task:", counter)
}

func does_contain(slice []int64, x int64) bool {
	for i := range slice {
		slice_element := slice[i]

		if slice_element == x {
			return true
		}
	}

	return false
}

func run_second() {
	file, err := os.Open("./input/4.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	counter := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitted_line := strings.Split(line, ",")
		left_range := make([]int64, 0)
		right_range := make([]int64, 0)

		for i := 0; i < 2; i++ {
			rangeString := splitted_line[i]

			begin, err := strconv.ParseInt(strings.Split(rangeString, "-")[0], 10, 16)
			check(err)
			end, err := strconv.ParseInt(strings.Split(rangeString, "-")[1], 10, 16)
			check(err)

			for j := begin; j <= end; j++ {
				if i == 0 {
					left_range = append(left_range, j)
				} else {
					right_range = append(right_range, j)
				}
			}

		}

		right_contains_left := false
		left_contains_right := false

		// Does right contain something of left?
		for i := range left_range {
			if does_contain(right_range, left_range[i]) {
				right_contains_left = true
				break
			}
		}

		// Does left contain something of right?
		for i := range right_range {
			if does_contain(left_range, right_range[i]) {
				left_contains_right = true
				break
			}
		}

		if left_contains_right || right_contains_left {
			counter++
		}

	}

	fmt.Println("Second task:", counter)
}
