package main

import (
	"fmt"
	"slices"

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

const (
	numPacks int = 12
)

func part2() {
	var total int
	utils.ReadFile("./cmd/day03/input.txt", func(line string) {
		bank := make([]int, 0)
		for _, v := range line {
			bank = append(bank, utils.Atoi(string(v)))
		}
		jolts := 0
		for index := range numPacks {
			lastIdx := (len(bank) - (numPacks - (index + 1)))
			digit := slices.Max(bank[:lastIdx])
			bank = bank[slices.Index(bank, digit)+1:]
			jolts = (jolts * 10) + digit
		}
		total += jolts

	})

	fmt.Println("Part 2", total)
}

func main() {
	part1()
	part2()
}
