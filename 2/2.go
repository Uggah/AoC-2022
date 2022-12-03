package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// run_first()
	run_second()
}

func run_first() {
	file, err := os.Open("./input/2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitted_line := strings.Split(line, " ")

		score_shape_selected := 0
		score_won := 0

		if splitted_line[1] == "X" {
			// ROCK
			score_shape_selected = 1

			if splitted_line[0] == "C" {
				score_won = 6
			} else if splitted_line[0] == "A" {
				score_won = 3
			}
		} else if splitted_line[1] == "Y" {
			// PAPER
			score_shape_selected = 2

			if splitted_line[0] == "A" {
				score_won = 6
			} else if splitted_line[0] == "B" {
				score_won = 3
			}
		} else if splitted_line[1] == "Z" {
			// SCISSORS
			score_shape_selected = 3

			if splitted_line[0] == "B" {
				score_won = 6
			} else if splitted_line[0] == "C" {
				score_won = 3
			}
		}

		score += (score_shape_selected + score_won)

	}

	fmt.Print(score)
}

func run_second() {
	file, err := os.Open("./input/2.sample.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitted_line := strings.Split(line, " ")

		score_shape_selected := 0
		score_won := 0

		if splitted_line[0] == "A" {
			if splitted_line[1] == "X" {
				score_shape_selected = 3 // SCISSORS
				score_won = 0
			} else if splitted_line[1] == "Y" {
				score_shape_selected = 1 // ROCK
				score_won = 3
			} else if splitted_line[1] == "Z" {
				score_shape_selected = 2 // PAPER
				score_won = 6
			}
		} else if splitted_line[0] == "B" {
			if splitted_line[1] == "X" {
				score_shape_selected = 1 // ROCK
				score_won = 0
			} else if splitted_line[1] == "Y" {
				score_shape_selected = 2 // PAPER
				score_won = 3
			} else if splitted_line[1] == "Z" {
				score_shape_selected = 3 // SCISSORS
				score_won = 6
			}
		} else if splitted_line[0] == "C" {
			if splitted_line[1] == "X" {
				score_shape_selected = 2 // PAPER
				score_won = 0
			} else if splitted_line[1] == "Y" {
				score_shape_selected = 3 // SCISSORS
				score_won = 3
			} else if splitted_line[1] == "Z" {
				score_shape_selected = 1 // ROCK
				score_won = 6
			}
		}

		score += (score_shape_selected + score_won)
	}

	fmt.Print(score)
}
