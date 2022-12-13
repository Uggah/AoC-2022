package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
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
	path := "./input/13.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	index := 0
	sum := 0

	for scanner.Scan() {
		first_line := scanner.Text()

		if first_line != "" {
			index++
			scanner.Scan()
			second_line := scanner.Text()

			right_order := compare(first_line, second_line)

			if right_order == 1 {
				sum += index
			}
		}
	}

	fmt.Println("First task:", sum)
}

func run_second() {
	path := "./input/13.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	list := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			list = append(list, line)
		}
	}

	list = append(list, "[[2]]", "[[6]]")

	sort.Slice(list, func(i, j int) bool {
		return compare(list[i], list[j]) == 1
	})

	decoder_key := 1

	for i, element := range list {
		if element == "[[2]]" || element == "[[6]]" {
			decoder_key *= i + 1
		}
	}

	fmt.Println("Second task:", decoder_key)
}

func compare(left any, right any) int {
	if typeof(left) == "[]interface {}" && typeof(right) == "float64" {
		right = []any{right}
	}

	if typeof(left) == "float64" && typeof(right) == "[]interface {}" {
		left = []any{left}
	}

	if typeof(left) == "string" && typeof(right) == "string" {
		left_slice := parse(left.(string))
		right_slice := parse(right.(string))

		llen, rlen := len(left_slice), len(right_slice)

		for i := 0; i < get_minimum(llen, rlen); i++ {
			result := compare(left_slice[i], right_slice[i])

			if result != 0 {
				return result
			}
		}

		if llen < rlen {
			return 1
		} else if llen > rlen {
			return -1
		} else {
			return 0
		}
	}

	if typeof(left) == "float64" && typeof(right) == "float64" {
		if left.(float64) < right.(float64) {
			return 1
		}

		if left.(float64) == right.(float64) {
			return 0
		}

		if left.(float64) > right.(float64) {
			return -1
		}
	}

	if typeof(left) == "[]interface {}" && typeof(right) == "[]interface {}" {
		left_slice := left.([]interface{})
		right_slice := right.([]interface{})

		llen, rlen := len(left_slice), len(right_slice)

		for i := 0; i < get_minimum(llen, rlen); i++ {
			result := compare(left_slice[i], right_slice[i])

			if result != 0 {
				return result
			}
		}

		if llen < rlen {
			return 1
		} else if llen > rlen {
			return -1
		} else {
			return 0
		}
	}

	return -1
}

func parse(input string) []any {
	var returnList []any
	json.Unmarshal([]byte(input), &returnList)

	return returnList
}

func typeof(input any) string {
	return reflect.TypeOf(input).String()
}

func get_minimum(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
