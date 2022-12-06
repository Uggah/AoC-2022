package main

import (
	"bufio"
	"fmt"
	"os"
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
	path := "./input/6.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	scanner.Scan()

	line := scanner.Text()

	last_four_chars := make([]byte, 0)
	counter := 0

	for i := range line {
		char := line[i]

		if len(last_four_chars) == 4 {
			if chars_unique(last_four_chars) {
				break
			}

			last_four_chars = last_four_chars[1:4]
		}

		last_four_chars = append(last_four_chars, char)

		counter++

	}

	fmt.Println("First task:", counter)
}

func run_second() {
	path := "./input/6.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	scanner.Scan()

	line := scanner.Text()

	last_fourteen_chars := make([]byte, 0)
	counter := 0

	for i := range line {
		char := line[i]

		if len(last_fourteen_chars) == 14 {
			if chars_unique(last_fourteen_chars) {
				break
			}

			last_fourteen_chars = last_fourteen_chars[1:14]
		}

		last_fourteen_chars = append(last_fourteen_chars, char)

		counter++

	}

	fmt.Println("Second task:", counter)

}

// HELPER FUNCTIONS

func chars_unique(chars []byte) bool {
	for i := range chars {
		char := chars[i]

		for j := i + 1; j < len(chars); j++ {
			if char == chars[j] {
				return false
			}
		}
	}

	return true
}
