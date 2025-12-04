package main

import (
	"fmt"
	"math"

	"github.com/umbe77/aoc-2025/utils"
)

func countNeighbors(y int, x int, storage [][]string) int {
	neighbors := 0

	for dy := y - 1; dy <= y+1; dy++ {
		if dy < 0 || dy >= len(storage) {
			continue
		}
		for dx := x - 1; dx <= x+1; dx++ {
			if dx < 0 || dx >= len(storage[dy]) {
				continue
			}

			if dx == x && dy == y {
				continue
			}

			if storage[dy][dx] == "@" {
				neighbors++
			}
		}
	}

	return neighbors
}

func part1() {
	var count int
	storage := make([][]string, 0)
	utils.ReadFile("./cmd/day04/input.txt", func(line string) {
		cols := make([]string, 0)
		for _, c := range line {
			cols = append(cols, string(c))
		}
		storage = append(storage, cols)
	})

	for y := 0; y < len(storage); y++ {
		line := storage[y]
		for x := range line {
			if line[x] == "@" {
				neighbors := countNeighbors(y, x, storage)
				if neighbors < 4 {
					count++
				}
			}

		}
	}

	fmt.Println("Part 1:", count)
}

func part2() {
	var count int
	storage := make([][]string, 0)
	nextStorage := make([][]string, 0)

	utils.ReadFile("./cmd/day04/input.txt", func(line string) {
		cols := make([]string, 0)
		for _, c := range line {
			cols = append(cols, string(c))
		}
		storage = append(storage, cols)
	})

	currentCount := math.MaxInt
	for currentCount > 0 {
		currentCount = 0
		for y := 0; y < len(storage); y++ {
			line := storage[y]
			nextStorage = append(nextStorage, line)
			for x := range line {
				nextStorage[y][x] = line[x]
				if line[x] == "@" {
					neighbors := countNeighbors(y, x, storage)
					if neighbors < 4 {
						currentCount++
						nextStorage[y][x] = "."
					}
				}
			}
		}
		count += currentCount
		storage = nextStorage
	}

	fmt.Println("Part 2", count)
}

func main() {
	part1()
	part2()
}
