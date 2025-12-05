package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/umbe77/aoc-2025/utils"
)

type FreshIdsRange struct {
	start int
	end   int
}

func part1() {
	ranges := make([]FreshIdsRange, 0)
	ids := make([]int, 0)
	rangeProcessing := true
	utils.ReadFile("./cmd/day05/input.txt", func(line string) {
		if line == "" {
			rangeProcessing = false
			return
		}
		if rangeProcessing {
			parts := strings.Split(line, "-")
			ranges = append(ranges, FreshIdsRange{
				start: utils.Atoi(parts[0]),
				end:   utils.Atoi(parts[1]),
			})
		}
		if !rangeProcessing {
			ids = append(ids, utils.Atoi(line))
		}
	})

	fresh := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				fresh++
				break
			}
		}
	}

	fmt.Println("Part 1:", fresh)
}

func part2() {
	ranges := make([]FreshIdsRange, 0)
	ids := make([]int, 0)
	rangeProcessing := true
	utils.ReadFile("./cmd/day05/input.txt", func(line string) {
		if line == "" {
			rangeProcessing = false
			return
		}
		if rangeProcessing {
			parts := strings.Split(line, "-")
			ranges = append(ranges, FreshIdsRange{
				start: utils.Atoi(parts[0]),
				end:   utils.Atoi(parts[1]),
			})
		}
		if !rangeProcessing {
			ids = append(ids, utils.Atoi(line))
		}
	})

	slices.SortFunc(ranges, func(a FreshIdsRange, b FreshIdsRange) int {
		if a.start < b.start {
			return -1
		}
		if a.start > b.start {
			return 1
		}
		return 0
	})

	fresh := 0
	previous := FreshIdsRange{end: 0}
	for _, r := range ranges {

		if r.start > previous.end {
			fresh += r.end - r.start + 1
			previous = r
			continue
		}

		if r.start <= previous.end && r.end > previous.end {
			fresh += r.end - previous.end
			previous = r
			continue
		}

	}
	fmt.Println("Part 2", fresh)
}

func main() {
	part1()
	part2()
}
