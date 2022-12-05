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
	path := "./input/5.txt"

	stacks := read_stacks(path)

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line[0] != 'm' {
			continue
		}

		amount, from, to := read_instruction(line)

		fmt.Printf("Instruction: Move %d from %d to %d. \n", amount, from, to)

		fmt.Println("Stacks before execution:", log_stacks(stacks))

		from_stack := stacks[from-1]
		from_stack_size := len(from_stack)

		for i := 0; int64(i) < amount; i++ {
			stacks[to-1] = append(stacks[to-1], from_stack[from_stack_size-i-1])
			stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
		}

		fmt.Println("Stacks after execution:", log_stacks(stacks))
	}

	fmt.Println("First task:")

	for i := range stacks {
		stack := stacks[i]
		fmt.Print(string(stack[len(stack)-1]))
	}
}

func run_second() {
	path := "./input/5.txt"

	stacks := read_stacks(path)

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line[0] != 'm' {
			continue
		}

		amount, from, to := read_instruction(line)

		fmt.Printf("Instruction: Move %d from %d to %d. \n", amount, from, to)

		fmt.Println("Stacks before execution:", log_stacks(stacks))

		from_stack := stacks[from-1]
		from_stack_size := int64(len(from_stack))

		from_elements := from_stack[(from_stack_size - amount):from_stack_size]

		stacks[from-1] = stacks[from-1][0:(from_stack_size - amount)]

		stacks[to-1] = append(stacks[to-1], from_elements...)

		fmt.Println("Stacks after execution:", log_stacks(stacks))
	}

	fmt.Println("Second task:")

	for i := range stacks {
		stack := stacks[i]
		fmt.Print(string(stack[len(stack)-1]))
	}

}

// HELPER FUNCTIONS

func get_stack_count(path string) int {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	return (len(line) + 1) / 4
}

func read_stacks(path string) [][]byte {
	stack_count := get_stack_count(path)

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stacks := make([][]byte, stack_count)

	for i := range stacks {
		stacks[i] = make([]byte, 0)
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		for i := 0; i < stack_count*4; i += 4 {
			stack := stacks[i/4]

			letter := line[i+1]

			if letter >= 65 && letter <= 90 {
				stacks[i/4] = append([]byte{letter}, stack...)
			}

		}

	}

	return stacks
}

func read_instruction(line string) (int64, int64, int64) {
	var err error
	var amount int64
	var offset int

	if line[6] == ' ' {
		amount, err = strconv.ParseInt(string(line[5]), 10, 64)
		check(err)

		offset = 0
	} else {
		amount, err = strconv.ParseInt(string(line[5:7]), 10, 64)
		check(err)

		offset = 1
	}

	from, err := strconv.ParseInt(string(line[12+offset]), 10, 64)
	check(err)

	to, err := strconv.ParseInt(string(line[17+offset]), 10, 64)
	check(err)

	return amount, from, to
}

func log_stack(stack []byte) string {
	returnString := ""

	for i := range stack {
		returnString = fmt.Sprint(returnString, "[", string(stack[i]), "]")
	}

	return returnString
}

func log_stacks(stacks [][]byte) string {
	returnString := "\n"

	for i := range stacks {
		returnString = fmt.Sprint(returnString, i, ": ", log_stack(stacks[i]), "\n")
	}

	return returnString
}
