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
	path := "./input/7.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	file_map := make(map[string]int64)

	dir_name := ""
	var accumulated_file_size int64
	accumulated_file_size = 0
	contained_dirs := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == '$' {
			if line[:4] == "$ cd" {
				path := line[5:]

				if path == "/" {
					dir_name = "/"
					continue
				}

				if _, exists := file_map[dir_name]; !exists {
					file_map[dir_name] = accumulated_file_size
				}

				if path == ".." {
					dir_structure := strings.Split(dir_name, "/")
					dir_name = strings.Join(dir_structure[:len(dir_structure)-2], "/") + "/"
					accumulated_file_size = 0
				} else {
					dir_name += path + "/"
					accumulated_file_size = 0
				}
			}
			continue
		}

		if line[:3] != "dir" {
			file_size, err := strconv.ParseInt(strings.Split(line, " ")[0], 10, 64)
			check(err)
			accumulated_file_size += file_size

			file_map[dir_name] = accumulated_file_size
		} else {
			if dir_name != "" {
				contained_dirs[dir_name] = append(contained_dirs[dir_name], line[4:])
			}
		}
	}

	var counter int64
	counter = 0

	for k, _ := range file_map {

		true_size := get_dir_size(k, file_map)

		if k != "" {
			if true_size <= 100000 {
				counter += true_size
			}

		}
	}

	fmt.Println("First task:", counter)
}

func run_second() {
	path := "./input/7.txt"

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)

	file_map := make(map[string]int64)

	dir_name := ""
	var accumulated_file_size int64
	accumulated_file_size = 0
	contained_dirs := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == '$' {
			if line[:4] == "$ cd" {
				path := line[5:]

				if path == "/" {
					dir_name = "/"
					continue
				}

				if _, exists := file_map[dir_name]; !exists {
					file_map[dir_name] = accumulated_file_size
				}

				if path == ".." {
					dir_structure := strings.Split(dir_name, "/")
					dir_name = strings.Join(dir_structure[:len(dir_structure)-2], "/") + "/"
					accumulated_file_size = 0
				} else {
					dir_name += path + "/"
					accumulated_file_size = 0
				}
			}
			continue
		}

		if line[:3] != "dir" {
			file_size, err := strconv.ParseInt(strings.Split(line, " ")[0], 10, 64)
			check(err)
			accumulated_file_size += file_size

			file_map[dir_name] = accumulated_file_size
		} else {
			if dir_name != "" {
				contained_dirs[dir_name] = append(contained_dirs[dir_name], line[4:])
			}
		}
	}

	root_size := get_dir_size("/", file_map)
	needed_size := 30000000 - (70000000 - root_size)

	var min int64
	min = root_size

	for k, _ := range file_map {
		true_size := get_dir_size(k, file_map)

		if true_size >= needed_size && true_size < min {
			min = true_size
		}
	}

	fmt.Println("Second task", min)
}

func get_dir_size(dir string, file_map map[string]int64) int64 {
	length := len(dir)

	var size int64
	size = 0
	for k, v := range file_map {
		if len(k) >= length && k[:length] == dir {
			size += v
		}
	}

	return size
}
