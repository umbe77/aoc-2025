package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/umbe77/aoc-2025/utils"
)

func part1() {
	movements := make([]int, 0)
	utils.ReadFile("./cmd/day01/input.txt", func(line string) {
		a := []rune(line)
		direction := string(a[0])
		steps, err := strconv.Atoi(strings.TrimSpace(string(a[1:])))
		if err != nil {
			panic(err)
		}
		if direction == "L" {
			steps = -steps
		}
		movements = append(movements, steps)
	})

	pwd := 0
	currentPos := 50
	for _, move := range movements {
		if move > 0 {
			currentPos = (currentPos + move) % 100

		} else {
			currentPos = currentPos - (-move % 100)
			if currentPos < 0 {
				currentPos = 100 + currentPos
			}
		}
		if currentPos == 0 {
			pwd = pwd + 1
		}
	}

	fmt.Println("Part 1:", pwd)
}

func countZeroTouches(start, amount, size int, direction byte) (end int, touches int) {

	if direction == 'R' {
		// Final position
		end = (start + amount) % size

		// Number of times we wrap past zero
		touches = (start + amount) / size
		return
	}

	// LEFT rotation:
	// Final position
	end = (start - amount%size + size) % size

	// Number of wraps crossing zero
	touches = max((amount-start+size-1)/size, 0)
	return
}

func part2() {
	movements := make([]int, 0)
	utils.ReadFile("./cmd/day01/input.txt", func(line string) {
		a := []rune(line)
		direction := string(a[0])
		steps, err := strconv.Atoi(strings.TrimSpace(string(a[1:])))
		if err != nil {
			panic(err)
		}
		if direction == "L" {
			steps = -steps
		}
		movements = append(movements, steps)
	})

	const size = 100
	position := 50
	totalTouches := 0

	for _, s := range movements {
		dir := byte('R')
		if s < 0 {
			dir = 'L'
		}
		amount := utils.Abs(s)
		newPos, touches := countZeroTouches(position, amount, size, dir)
		totalTouches += touches
		position = newPos
	}

	// totalTouches := 0
	// currentPos := 50
	// for _, move := range movements {
	// 	if move > 0 {
	// 		step := currentPos + move
	// 		currentPos = (step) % 100
	// 		if step >= 100 {
	// 			pwd += (step / 100)
	// 		}
	//
	// 	} else {
	//
	// 		steps := utils.Abs(move)
	// 		diff := currentPos - steps
	// 		if diff < 0 {
	// 			step := math.Ceil(math.Abs(float64(diff)) / 100)
	// 			pwd += int(step)
	// 		}
	// 		currentPos = ((diff % 100) + 100) % 100
	//
	// 	}
	// }
	//
	fmt.Println("Part 2:", totalTouches)
}

func main() {
	part1()
	part2()
}
