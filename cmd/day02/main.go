package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/umbe77/aoc-2025/utils"
)

type Range struct {
	Start int
	End   int
}

func part1() {
	// Implementation for part 1
	ranges := make([]Range, 0)
	utils.ReadFile("./cmd/day02/input.txt", func(line string) {
		rngs := strings.SplitSeq(line, ",")
		for r := range rngs {
			bounds := strings.Split(r, "-")
			start, err := strconv.Atoi(bounds[0])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(bounds[1])
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, Range{Start: start, End: end})
		}
	})

	invalidNumbers := make([]int, 0)
	sum := 0

	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			c := strconv.Itoa(i)
			if (len(c) % 2) == 0 {
				a := c[:len(c)/2]
				b := c[len(c)/2:]
				if a == b {
					invalidNumbers = append(invalidNumbers, i)
				}
			}
		}
	}

	for _, v := range invalidNumbers {
		sum += v
	}

	fmt.Println("Part 1:", sum)
}

func part2() {
	// Implementation for part 2
}

func main() {
	part1()
	part2()
}
