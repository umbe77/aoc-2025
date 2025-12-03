package main

import (
	"fmt"

	"github.com/umbe77/aoc-2025/utils"
)

func part1() {
	var total int
	utils.ReadFile("./cmd/day03/input.txt", func(line string) {
		maxVal := 0
		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				n := utils.Atoi(string(line[i]) + string(line[j]))
				maxVal = utils.Max(maxVal, n)
			}
		}
		total += maxVal
	})

	fmt.Println("Part 1:", total)
}

func part2() {
	// utils.ReadFile("./cmd/day03/sample.txt", func(line string) {
	// })
}

func main() {
	part1()
	part2()
}
