package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Rock struct {
	x int
	y int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//run_first()
	run_second()
}

func run_first() {
	path := "./input/14.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	rocks := make([]Rock, 0)

	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, " -> ")

		curr_rocks := make([]Rock, 0)

		for i := range splitted {
			coords := strings.Split(splitted[i], ",")

			x, err := strconv.ParseInt(coords[0], 10, 64)
			check(err)
			y, err := strconv.ParseInt(coords[1], 10, 64)
			check(err)

			if len(curr_rocks) > 0 {
				last_rock := curr_rocks[len(curr_rocks)-1]

				if last_rock.x != int(x) {
					if last_rock.x < int(x) {
						for x_1 := last_rock.x + 1; int64(x_1) <= x; x_1++ {
							curr_rocks = append(curr_rocks, Rock{x_1, int(y)})
						}
					} else {
						for x_1 := last_rock.x - 1; int64(x_1) >= x; x_1-- {
							curr_rocks = append(curr_rocks, Rock{x_1, int(y)})
						}
					}
				} else {
					if last_rock.y < int(y) {
						for y_1 := last_rock.y + 1; int64(y_1) <= y; y_1++ {
							curr_rocks = append(curr_rocks, Rock{int(x), y_1})
						}
					} else {
						for y_1 := last_rock.y - 1; int64(y_1) >= y; y_1-- {
							curr_rocks = append(curr_rocks, Rock{int(x), y_1})
						}
					}
				}

			} else {
				curr_rocks = append(curr_rocks, Rock{int(x), int(y)})
			}
		}

		rocks = append(rocks, curr_rocks...)

	}

	out := false
	sand := make([]Rock, 0)
	counter := 0

	_, _, max_y, _ := get_min_max(rocks)

	for !out {
		curr_sand := Rock{500, 0}

		for {
			//...
			//.o.
			//###
			if exists(curr_sand.x, curr_sand.y+1, rocks, sand) && exists(curr_sand.x-1, curr_sand.y+1, rocks, sand) && exists(curr_sand.x+1, curr_sand.y+1, rocks, sand) {
				counter++
				break
			}

			//...
			//.o.
			//##.
			if exists(curr_sand.x, curr_sand.y+1, rocks, sand) && exists(curr_sand.x-1, curr_sand.y+1, rocks, sand) {
				curr_sand = Rock{curr_sand.x + 1, curr_sand.y + 1}
				continue
			}

			//...
			//.o.
			//.#.
			if exists(curr_sand.x, curr_sand.y+1, rocks, sand) {
				curr_sand = Rock{curr_sand.x - 1, curr_sand.y + 1}
				continue
			}

			if curr_sand.y == max_y {
				out = true
				break
			}

			//...
			//.o.
			//...
			curr_sand = Rock{curr_sand.x, curr_sand.y + 1}

		}

		sand = append(sand, curr_sand)
	}

	fmt.Println("First task:", counter)
	print(rocks)
}

func run_second() {
	path := "./input/14.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	rocks := make([]Rock, 0)

	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, " -> ")

		curr_rocks := make([]Rock, 0)

		for i := range splitted {
			coords := strings.Split(splitted[i], ",")

			x, err := strconv.ParseInt(coords[0], 10, 64)
			check(err)
			y, err := strconv.ParseInt(coords[1], 10, 64)
			check(err)

			if len(curr_rocks) > 0 {
				last_rock := curr_rocks[len(curr_rocks)-1]

				if last_rock.x != int(x) {
					if last_rock.x < int(x) {
						for x_1 := last_rock.x + 1; int64(x_1) <= x; x_1++ {
							curr_rocks = append(curr_rocks, Rock{x_1, int(y)})
						}
					} else {
						for x_1 := last_rock.x - 1; int64(x_1) >= x; x_1-- {
							curr_rocks = append(curr_rocks, Rock{x_1, int(y)})
						}
					}
				} else {
					if last_rock.y < int(y) {
						for y_1 := last_rock.y + 1; int64(y_1) <= y; y_1++ {
							curr_rocks = append(curr_rocks, Rock{int(x), y_1})
						}
					} else {
						for y_1 := last_rock.y - 1; int64(y_1) >= y; y_1-- {
							curr_rocks = append(curr_rocks, Rock{int(x), y_1})
						}
					}
				}

			} else {
				curr_rocks = append(curr_rocks, Rock{int(x), int(y)})
			}
		}

		rocks = append(rocks, curr_rocks...)

	}

	_, _, max_y, _ := get_min_max(rocks)

	out := false
	sand := make([]Rock, 0)
	counter := 0

	for !out {
		curr_sand := Rock{500, 0}

		for {
			if curr_sand.y == max_y+2 {
				break
			}

			//...
			//.o.
			//###
			if exists(curr_sand.x, curr_sand.y+1, rocks, sand) && exists(curr_sand.x-1, curr_sand.y+1, rocks, sand) && exists(curr_sand.x+1, curr_sand.y+1, rocks, sand) {
				counter++
				if curr_sand.y == 0 {
					out = true
				}
				break
			}

			//...
			//.o.
			//##.
			if exists(curr_sand.x, curr_sand.y+1, rocks, sand) && exists(curr_sand.x-1, curr_sand.y+1, rocks, sand) {
				curr_sand.x++
				curr_sand.y++
				continue
			}

			//...
			//.o.
			//.#.
			if exists(curr_sand.x, curr_sand.y+1, rocks, sand) {
				curr_sand.x--
				curr_sand.y++
				continue
			}

			//...
			//.o.
			//...
			curr_sand.y++

		}

		sand = append(sand, curr_sand)
	}

	fmt.Println("Second task:", counter)
	print(rocks)
}

func get_min_max(rocks []Rock) (int, int, int, int) {
	max_x := 0
	min_x := math.MaxInt
	max_y := 0
	min_y := math.MaxInt

	for _, rock := range rocks {
		if rock.x > max_x {
			max_x = rock.x
		}

		if rock.x < min_x {
			min_x = rock.x
		}

		if rock.y > max_y {
			max_y = rock.y
		}

		if rock.y < min_y {
			min_y = rock.y
		}
	}

	return max_x, min_x, max_y, min_y

}

func print(rocks []Rock) {
	max_x, min_x, max_y, _ := get_min_max(rocks)

	rock_map := make([][]bool, max_y+1)

	for i := range rock_map {
		rock_map[i] = make([]bool, max_x+1)
	}

	for _, rock := range rocks {
		rock_map[rock.y][rock.x] = true
	}

	for y := 0; y <= max_y; y++ {
		fmt.Print(y, ": ")
		for x := min_x; x <= max_x; x++ {
			if rock_map[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

}

func exists(x int, y int, rocks []Rock, sand []Rock) bool {
	for _, element := range rocks {
		if element.x == x && element.y == y {
			return true
		}
	}

	for _, element := range sand {
		if element.x == x && element.y == y {
			return true
		}
	}

	return false
}
