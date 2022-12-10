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
	path := "./input/10.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	scheduled := false

	cycle := 1
	x := 1

	instructions := make([]string, 0)
	executions := 0

	signal_strength_sum := 0

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	for executions < len(instructions) {
		instruction := instructions[executions]

		if cycle == 20 || (cycle-20)%40 == 0 {
			signal_strength_sum += cycle * x
		}

		if instruction == "noop" {
			cycle++
			executions++
			continue
		}

		if !scheduled {
			scheduled = true
		} else {
			splitted_instruction := strings.Split(instruction, " ")
			parsed_int, err := strconv.ParseInt(splitted_instruction[1], 10, 64)
			check(err)

			x += int(parsed_int)
			executions++
			scheduled = false
		}

		cycle++

	}

	fmt.Println("First task:", signal_strength_sum)
}

func run_second() {
	path := "./input/10.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	screen := make([][]bool, 6)

	for i := range screen {
		screen[i] = make([]bool, 40)
	}

	curr_sprite := "###....................................."

	scheduled := false

	cycle := 1
	x := 1

	instructions := make([]string, 0)
	executions := 0

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	for executions < len(instructions) {
		instruction := instructions[executions]

		screen_y := (cycle - 1) / 40
		screen_x := (cycle - 1) % 40

		screen[screen_y][screen_x] = curr_sprite[screen_x] == '#'

		if instruction == "noop" {
			cycle++
			executions++
			continue
		}

		if !scheduled {
			scheduled = true
		} else {
			splitted_instruction := strings.Split(instruction, " ")
			parsed_int, err := strconv.ParseInt(splitted_instruction[1], 10, 64)
			check(err)

			x += int(parsed_int)

			curr_sprite = draw_sprite(int(x))

			executions++
			scheduled = false
		}

		cycle++

	}

	fmt.Println("Second task:")
	draw_screen(screen)
}

func draw_sprite(pos int) string {
	sprite := make([]byte, 40)

	for i := range sprite {
		if i == pos-1 || i == pos || i == pos+1 {
			sprite[i] = '#'
		} else {
			sprite[i] = '.'
		}
	}

	return string(sprite)
}

func draw_screen(screen [][]bool) {
	for i := range screen {
		for j := range screen[i] {
			if screen[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}
}
