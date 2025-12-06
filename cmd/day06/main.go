package main

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2025/utils"
)

type worksheetcol struct {
	values    []int
	operation string
}

type worksheet struct {
	cols []worksheetcol
}

func part1() {

	ws := worksheet{}
	columnCount := 0
	utils.ReadFile("./cmd/day06/input.txt", func(line string) {
		c := strings.Split(line, " ")
		cols := make([]string, 0)
		for _, col := range c {
			v := strings.Trim(col, " ")
			if v != "" {
				cols = append(cols, v)
			}
		}

		if columnCount == 0 {
			columnCount = len(cols)
			ws.cols = make([]worksheetcol, columnCount)
			for i := range columnCount {
				ws.cols[i] = worksheetcol{
					values: make([]int, 0),
				}
			}
		}

		if cols[0] == "+" || cols[0] == "*" {
			// Operators
			for i, op := range cols {
				ws.cols[i].operation = op
			}
			return
		}

		for i, col := range cols {
			v := utils.Atoi(col)
			ws.cols[i].values = append(ws.cols[i].values, v)
		}

	})

	total := 0
	for _, col := range ws.cols {
		colTotal := 0
		if col.operation == "*" {
			colTotal = 1
		}
		for _, v := range col.values {
			switch col.operation {
			case "+":
				colTotal += v
			case "*":
				colTotal *= v
			}
		}
		total += colTotal
	}

	fmt.Println("Part 1:", total)
}

func part2() {
	ws := worksheet{
		cols: make([]worksheetcol, 0),
	}
	lines := make([]string, 0)

	utils.ReadFile("./cmd/day06/input.txt", func(line string) {
		lines = append(lines, line)
	})

	col := worksheetcol{
		values: make([]int, 0),
	}

	for c := len(lines[0]) - 1; c >= 0; c-- {
		n := ""
		for i, l := range lines {
			if i == len(lines)-1 {
				col.values = append(col.values, utils.Atoi(n))

				if l[c] != ' ' {
					col.operation = string(l[c])
					ws.cols = append(ws.cols, col)
					col = worksheetcol{
						values: make([]int, 0),
					}
					c--
					break
				}
			}
			if l[c] != ' ' {
				n += string(l[c])
			}
		}
	}

	total := 0
	for _, col := range ws.cols {
		colTotal := 0
		if col.operation == "*" {
			colTotal = 1
		}
		for _, v := range col.values {
			switch col.operation {
			case "+":
				colTotal += v
			case "*":
				colTotal *= v
			}
		}
		total += colTotal
	}

	fmt.Println("Part 2:", total)
}

func main() {
	part1()
	part2()
}
