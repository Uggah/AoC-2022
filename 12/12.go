package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	path := "./input/12.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	elevation := make([][]int, 0)

	source := 0
	destination := 0

	for scanner.Scan() {
		line := scanner.Text()
		current_row := make([]int, 0)

		for i := range line {
			current_elevation := line[i]

			if current_elevation == 'S' {
				source = get_position_by_coordinates(i, len(elevation), len(line))
			} else if current_elevation == 'E' {
				destination = get_position_by_coordinates(i, len(elevation), len(line))
			}

			current_elevation_int := get_elevation_from_char(current_elevation)
			current_row = append(current_row, current_elevation_int)
		}

		elevation = append(elevation, current_row)

	}

	//print_elevation(elevation)

	adjacency_list := create_adjacency_list(elevation)

	fmt.Println("First task:", bfs(adjacency_list, source, destination))
}

func run_second() {
	path := "./input/12.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	elevation := make([][]int, 0)

	destination := 0

	for scanner.Scan() {
		line := scanner.Text()
		current_row := make([]int, 0)

		for i := range line {
			current_elevation := line[i]

			if current_elevation == 'E' {
				destination = get_position_by_coordinates(i, len(elevation), len(line))
			}

			current_elevation_int := get_elevation_from_char(current_elevation)
			current_row = append(current_row, current_elevation_int)
		}

		elevation = append(elevation, current_row)

	}

	//print_elevation(elevation)

	adjacency_list := create_adjacency_list(elevation)

	distances := make([]int, 0)

	for y := range elevation {
		for x := range elevation[y] {
			if elevation[y][x] == 0 {
				position := get_position_by_coordinates(x, y, len(elevation[y]))
				distance := bfs(adjacency_list, position, destination)

				if distance != 0 {
					distances = append(distances, distance)
				}
			}
		}
	}

	sort.Ints(distances)

	fmt.Println("Second task:", distances[0])
}

func get_elevation_from_char(char byte) int {
	if char == 'S' {
		return 0
	} else if char == 'E' {
		return 25
	}

	return int(char - 97)
}

func create_adjacency_list(elevation [][]int) [][]int {
	returnMatrix := make([][]int, len(elevation)*len(elevation[0]))

	for i := range returnMatrix {
		returnMatrix[i] = make([]int, 0)
	}

	for y := range elevation {
		for x := range elevation[y] {
			current_elevation := elevation[y][x] + 1
			current_position := get_position_by_coordinates(x, y, len(elevation[0]))
			//###
			//ba#
			//###
			if x > 0 && elevation[y][x-1] <= current_elevation {
				nextPosition := get_position_by_coordinates(x-1, y, len(elevation[0]))
				returnMatrix[current_position] = append(returnMatrix[current_position], nextPosition)
			}

			//###
			//#ab
			//###
			if x < len(elevation[y])-1 && elevation[y][x+1] <= current_elevation {
				nextPosition := get_position_by_coordinates(x+1, y, len(elevation[0]))
				returnMatrix[current_position] = append(returnMatrix[current_position], nextPosition)
			}

			//#b#
			//#a#
			//###
			if y > 0 && elevation[y-1][x] <= current_elevation {
				nextPosition := get_position_by_coordinates(x, y-1, len(elevation[0]))
				returnMatrix[current_position] = append(returnMatrix[current_position], nextPosition)
			}

			//###
			//#a#
			//#b#
			if y < len(elevation)-1 && elevation[y+1][x] <= current_elevation {
				nextPosition := get_position_by_coordinates(x, y+1, len(elevation[0]))
				returnMatrix[current_position] = append(returnMatrix[current_position], nextPosition)
			}
		}
	}

	return returnMatrix
}

func get_position_by_coordinates(x int, y int, len int) int {
	return len*(y) + x
}

func bfs(adjacent [][]int, source int, destination int) int {
	queue := make([]int, 0)

	visited := make([]bool, len(adjacent))
	predecessors := make([]int, len(adjacent))
	distances := make([]int, len(adjacent))

	for i := 0; i < len(adjacent); i++ {
		visited[i] = false
		distances[i] = math.MaxInt
		predecessors[i] = -1
	}

	visited[source] = true
	distances[source] = 0
	queue = append(queue, source)

	for len(queue) != 0 {
		element := queue[0]
		queue = queue[1:]

		for i := range adjacent[element] {
			if !visited[adjacent[element][i]] {
				visited[adjacent[element][i]] = true
				distances[adjacent[element][i]] = distances[element] + 1
				predecessors[adjacent[element][i]] = element
				queue = append(queue, adjacent[element][i])

				if adjacent[element][i] == destination {
					return distances[adjacent[element][i]]
				}
			}
		}
	}

	return 0
}

func print_elevation(elevation [][]int) {
	for y := range elevation {
		for x := range elevation[y] {
			fmt.Print(string(byte(elevation[y][x] + 97)))
		}

		fmt.Print("\n")
	}
}
