package main

import (
	"bufio"
	"fmt"
	"math/big"
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

type Monkey struct {
	items     []int
	operation string
	test      int
	ifTrue    int
	ifFalse   int
}

type Monkey2 struct {
	items     []big.Int
	operation string
	test      big.Int
	ifTrue    int
	ifFalse   int
}

func run_first() {
	path := "./input/11.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	monkeys := make([]Monkey, 0)

	curr_monkey_items := make([]int, 0)
	curr_monkey_operation := ""
	curr_monkey_division_test := 0
	curr_monkey_if_true := 0
	curr_monkey_if_false := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "  Starting items") {
			items_string := strings.Split(line, ": ")[1]
			numbers_string := strings.Split(items_string, ", ")

			for i := range numbers_string {
				parsed, err := strconv.ParseInt(numbers_string[i], 10, 64)
				check(err)

				curr_monkey_items = append(curr_monkey_items, int(parsed))
			}
		} else if strings.HasPrefix(line, "  Operation") {
			curr_monkey_operation = strings.Split(line, ": ")[1]
		} else if strings.HasPrefix(line, "  Test") {
			splitted := strings.Split(line, " ")
			parsed, err := strconv.ParseInt(splitted[len(splitted)-1], 10, 64)
			check(err)
			curr_monkey_division_test = int(parsed)
		} else if strings.HasPrefix(line, "    If true") {
			parsed, err := strconv.ParseInt(string(line[len(line)-1]), 10, 64)
			check(err)
			curr_monkey_if_true = int(parsed)
		} else if strings.HasPrefix(line, "    If false") {
			parsed, err := strconv.ParseInt(string(line[len(line)-1]), 10, 64)
			check(err)
			curr_monkey_if_false = int(parsed)

			monkeys = append(monkeys,
				Monkey{
					curr_monkey_items,
					curr_monkey_operation,
					curr_monkey_division_test,
					curr_monkey_if_true,
					curr_monkey_if_false,
				})

			curr_monkey_items = make([]int, 0)
			curr_monkey_operation = ""
			curr_monkey_division_test = 0
			curr_monkey_if_true = 0
			curr_monkey_if_false = 0
		}
	}

	inspections := make([]int, len(monkeys))

	for i := 0; i < 20; i++ {
		for j := range monkeys {
			monkey := monkeys[j]

			for k := range monkey.items {
				var result int
				op_type, first, second := get_operation_type(monkey, k)

				switch op_type {
				case '+':
					result = first + second
				case '-':
					result = first - second
				case '*':
					result = first * second
				case '/':
					result = first / second
				}

				result /= 3
				inspections[j] += 1

				if result%monkey.test == 0 {
					monkeys[monkey.ifTrue].items = append(monkeys[monkey.ifTrue].items, result)
				} else {
					monkeys[monkey.ifFalse].items = append(monkeys[monkey.ifFalse].items, result)
				}

				monkeys[j].items = monkeys[j].items[1:]
			}
		}
	}

	max_1 := 0
	max_2 := 0

	for i := range inspections {
		inspection := inspections[i]

		if inspection > max_1 {
			max_1 = inspection
		}
	}

	for i := range inspections {
		inspection := inspections[i]

		if inspection > max_2 && inspection != max_1 {
			max_2 = inspection
		}
	}

	fmt.Println("First task:", max_1*max_2)
}

func run_second() {
	path := "./input/11.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	monkeys := make([]Monkey, 0)

	curr_monkey_items := make([]int, 0)
	curr_monkey_operation := ""
	curr_monkey_division_test := 0
	curr_monkey_if_true := 0
	curr_monkey_if_false := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "  Starting items") {
			items_string := strings.Split(line, ": ")[1]
			numbers_string := strings.Split(items_string, ", ")

			for i := range numbers_string {
				parsed, err := strconv.ParseInt(numbers_string[i], 10, 64)
				check(err)

				curr_monkey_items = append(curr_monkey_items, int(parsed))
			}
		} else if strings.HasPrefix(line, "  Operation") {
			curr_monkey_operation = strings.Split(line, ": ")[1]
		} else if strings.HasPrefix(line, "  Test") {
			splitted := strings.Split(line, " ")
			parsed, err := strconv.ParseInt(splitted[len(splitted)-1], 10, 64)
			check(err)
			curr_monkey_division_test = int(parsed)
		} else if strings.HasPrefix(line, "    If true") {
			parsed, err := strconv.ParseInt(string(line[len(line)-1]), 10, 64)
			check(err)
			curr_monkey_if_true = int(parsed)
		} else if strings.HasPrefix(line, "    If false") {
			parsed, err := strconv.ParseInt(string(line[len(line)-1]), 10, 64)
			check(err)
			curr_monkey_if_false = int(parsed)

			monkeys = append(monkeys,
				Monkey{
					curr_monkey_items,
					curr_monkey_operation,
					curr_monkey_division_test,
					curr_monkey_if_true,
					curr_monkey_if_false,
				})

			curr_monkey_items = make([]int, 0)
			curr_monkey_operation = ""
			curr_monkey_division_test = 0
			curr_monkey_if_true = 0
			curr_monkey_if_false = 0
		}
	}

	inspections := make([]int, len(monkeys))

	for i := 0; i < 10000; i++ {
		for j := range monkeys {
			monkey := monkeys[j]

			for k := range monkey.items {
				var product int = 1
				for l := range monkeys {
					monkey2 := monkeys[l]

					product *= monkey2.test
				}

				var result int
				op_type, first, second := get_operation_type(monkey, k)

				switch op_type {
				case '+':
					result = first + second
				case '-':
					result = first - second
				case '*':
					result = first * second
				case '/':
					result = first / second
				}

				inspections[j] += 1

				result = result % product

				if result%monkey.test == 0 {
					monkeys[monkey.ifTrue].items = append(monkeys[monkey.ifTrue].items, result)
				} else {
					monkeys[monkey.ifFalse].items = append(monkeys[monkey.ifFalse].items, result)
				}

				monkeys[j].items = monkeys[j].items[1:]
			}
		}
	}

	max_1 := 0
	max_2 := 0

	for i := range inspections {
		inspection := inspections[i]

		if inspection > max_1 {
			max_1 = inspection
		}
	}

	for i := range inspections {
		inspection := inspections[i]

		if inspection > max_2 && inspection != max_1 {
			max_2 = inspection
		}
	}

	fmt.Println("Second task:", max_1*max_2)
}

func get_operation_type(monkey Monkey, item_index int) (byte, int, int) {
	operation := monkey.operation
	item := monkey.items[item_index]
	splitted := strings.Split(operation, " ")

	op_type := splitted[3][0]
	first := 0
	second := 0

	if splitted[2] == "old" {
		first = item
	} else {
		parsed, err := strconv.ParseInt(splitted[2], 10, 64)
		check(err)

		first = int(parsed)
	}

	if splitted[4] == "old" {
		second = item
	} else {
		parsed, err := strconv.ParseInt(splitted[4], 10, 64)
		check(err)

		second = int(parsed)
	}

	return op_type, first, second
}

// WARNING: Extremely inefficient!
func run_second_with_big_int() {
	path := "./input/11.sample.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	monkeys := make([]Monkey2, 0)

	curr_monkey_items := make([]big.Int, 0)
	curr_monkey_operation := ""
	var curr_monkey_division_test big.Int
	var curr_monkey_if_true int
	var curr_monkey_if_false int

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "  Starting items") {
			items_string := strings.Split(line, ": ")[1]
			numbers_string := strings.Split(items_string, ", ")

			for i := range numbers_string {
				parsed, err := strconv.ParseInt(numbers_string[i], 10, 64)
				check(err)

				curr_monkey_items = append(curr_monkey_items, *new(big.Int).SetInt64(parsed))
			}
		} else if strings.HasPrefix(line, "  Operation") {
			curr_monkey_operation = strings.Split(line, ": ")[1]
		} else if strings.HasPrefix(line, "  Test") {
			splitted := strings.Split(line, " ")
			parsed, err := strconv.ParseInt(splitted[len(splitted)-1], 10, 64)
			check(err)
			curr_monkey_division_test = *new(big.Int).SetInt64(parsed)
		} else if strings.HasPrefix(line, "    If true") {
			parsed, err := strconv.ParseInt(string(line[len(line)-1]), 10, 64)
			check(err)
			curr_monkey_if_true = int(parsed)
		} else if strings.HasPrefix(line, "    If false") {
			parsed, err := strconv.ParseInt(string(line[len(line)-1]), 10, 64)
			check(err)
			curr_monkey_if_false = int(parsed)

			monkeys = append(monkeys,
				Monkey2{
					curr_monkey_items,
					curr_monkey_operation,
					curr_monkey_division_test,
					curr_monkey_if_true,
					curr_monkey_if_false,
				})

			curr_monkey_items = make([]big.Int, 0)
			curr_monkey_operation = ""
			curr_monkey_division_test = *new(big.Int).SetInt64(0)
			curr_monkey_if_true = 0
			curr_monkey_if_false = 0
		}
	}

	inspections := make([]int, len(monkeys))

	for i := 0; i < 10000; i++ {
		for j := range monkeys {
			monkey := monkeys[j]

			for k := range monkey.items {
				var result big.Int
				op_type, first, second := get_operation_type_2(monkey, k)

				switch op_type {
				case '+':
					result = *new(big.Int).Add(&first, &second)
				case '-':
					result = *new(big.Int).Sub(&first, &second)
				case '*':
					result = *new(big.Int).Mul(&first, &second)
				case '/':
					result = *new(big.Int).Div(&first, &second)
				}

				inspections[j] += 1

				modulo := new(big.Int).Mod(&result, &monkey.test)
				if modulo.Cmp(new(big.Int).SetInt64(0)) == 0 {
					monkeys[monkey.ifTrue].items = append(monkeys[monkey.ifTrue].items, result)
				} else {
					monkeys[monkey.ifFalse].items = append(monkeys[monkey.ifFalse].items, result)
				}

				monkeys[j].items = monkeys[j].items[1:]
			}
		}

		fmt.Println("Inspections at round", i, ":", inspections)
	}

	max_1 := 0
	max_2 := 0

	for i := range inspections {
		inspection := inspections[i]

		if inspection > max_1 {
			max_1 = inspection
		}
	}

	for i := range inspections {
		inspection := inspections[i]

		if inspection > max_2 && inspection != max_1 {
			max_2 = inspection
		}
	}

	fmt.Println("Second task:", max_1*max_2)
}

func get_operation_type_2(monkey Monkey2, item_index int) (byte, big.Int, big.Int) {
	operation := monkey.operation
	item := monkey.items[item_index]
	splitted := strings.Split(operation, " ")

	op_type := splitted[3][0]
	var first big.Int
	var second big.Int

	if splitted[2] == "old" {
		first = item
	} else {
		first_pointer, ok := new(big.Int).SetString(splitted[2], 10)
		if !ok {
			panic(ok)
		}
		first = *first_pointer
	}

	if splitted[4] == "old" {
		second = item
	} else {
		second_pointer, ok := new(big.Int).SetString(splitted[4], 10)
		if !ok {
			panic(ok)
		}
		second = *second_pointer
	}

	return op_type, first, second
}
