package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("./input/1.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	counter := 0
	counterList := make([]int, 0)

	for scanner.Scan() {
		if scanner.Text() != "" {
			parsed, err := strconv.Atoi(scanner.Text())
			check(err)

			counter += parsed

			fmt.Println("Set counter to: ", counter)
		} else {
			fmt.Println("Resetting counter from: ", counter)

			counterList = append(counterList, counter)
			counter = 0
		}
	}

	counterList = append(counterList, counter)

	sort.Slice(counterList, func(i, j int) bool {
		return counterList[i] > counterList[j]
	})

	fmt.Println("Sorted list: ", counterList)
	fmt.Println("Sum of first three: ", counterList[0]+counterList[1]+counterList[2])

}
