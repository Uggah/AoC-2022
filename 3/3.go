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
	run_first()
	run_second()
}

func run_first() {
	file, err := os.Open("./input/3.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	doubleItems := make([]byte, 0)

	for scanner.Scan() {
		line := scanner.Text()
		first_compartment := line[:len(line)/2]
		second_compartment := line[len(line)/2:]

		for i := range first_compartment {
			first_letter := first_compartment[i]
			current_double_item := byte(0)

			for j := range second_compartment {
				second_letter := second_compartment[j]

				if first_letter == second_letter {
					current_double_item = first_letter
					break
				}
			}

			if current_double_item != 0 {
				doubleItems = append(doubleItems, current_double_item)
				break
			}
		}
	}

	priority_sum := 0

	for i := range doubleItems {
		priority := 0

		if doubleItems[i] >= 97 {
			priority = int(doubleItems[i] - 96)
		} else {
			priority = int(doubleItems[i] - 38)
		}

		priority_sum += priority
	}

	fmt.Println("Part one:", priority_sum)
}

func run_second() {
	file, err := os.Open("./input/3.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	lines := make([]string, 0)
	badges := make([]byte, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines)/3; i++ {
		rucksacks := lines[3*i : 3*i+3]

		for j := range rucksacks[0] {
			item := string(rucksacks[0][j])

			if strings.Contains(rucksacks[1], item) && strings.Contains(rucksacks[2], item) {
				badges = append(badges, item[0])
				break
			}
		}
	}

	priority_sum := 0

	for i := range badges {
		priority := 0

		if badges[i] >= 97 {
			priority = int(badges[i] - 96)
		} else {
			priority = int(badges[i] - 38)
		}

		priority_sum += priority
	}

	fmt.Println("Part two:", priority_sum)
}
