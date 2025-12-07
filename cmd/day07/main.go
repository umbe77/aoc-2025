package main

import (
	"fmt"

	"github.com/umbe77/aoc-2025/utils"
)

type Point struct {
	x int
	y int
	v string
}

func part1() {
	var previousState []bool
	// l := 0
	splitTimes := 0
	utils.ReadFile("./cmd/day07/input.txt", func(line string) {
		// l++
		if previousState == nil {
			previousState = make([]bool, len(line))
		}
		lineState := make([]bool, len(line))
		for i, v := range line {
			if v == 'S' {
				lineState[i] = true
				// fmt.Println("start", i)
				continue
			}
			if v == '.' && !lineState[i] {
				if previousState[i] {
					lineState[i] = true
					continue
				}
				lineState[i] = false
			}

			if v == '^' {
				// fmt.Println("split found", previousState[i])
				if previousState[i] {
					lineState[i-1] = true
					lineState[i+1] = true
					// fmt.Println(i, i-1, i+1, lineState[i-1], lineState[i+1])
					splitTimes++
				}
				lineState[i] = false
			}
		}
		// fmt.Println(l, lineState)
		previousState = lineState
	})

	fmt.Println("Part 1:", splitTimes)
}

var cache = make(map[Point]int)

func solve(grid [][]Point, p Point) int {
	if p.y >= len(grid) {
		return 1
	}

	if v, ok := cache[p]; ok {
		return v
	}

	if grid[p.y][p.x].v == "." || grid[p.y][p.x].v == "S" {
		if p.y+1 >= len(grid) {
			return 1
		}

		v := solve(grid, grid[p.y+1][p.x])
		cache[p] = v
		return v
	}

	if grid[p.y][p.x].v == "^" {
		v := solve(grid, grid[p.y][p.x-1]) + solve(grid, grid[p.y][p.x+1])
		cache[p] = v
		return v
	}
	return 0
}

func part2() {
	var start Point
	grid := make([][]Point, 0)
	y := 0
	utils.ReadFile("./cmd/day07/input.txt", func(line string) {
		l := make([]Point, 0)
		for x, c := range line {
			l = append(l, Point{x, y, string(c)})
			if c == 'S' {
				start = Point{x, y, "S"}
			}
		}
		grid = append(grid, l)
		y++
	})

	v := solve(grid, start)

	fmt.Println("Part 2", v)
}

func main() {
	part1()
	part2()
}
