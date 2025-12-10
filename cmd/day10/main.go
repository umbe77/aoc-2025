package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/umbe77/aoc-2025/utils"
)

type QueueItem struct {
	state   int
	presses int
}

type QueueItem2 struct {
	state   []int
	presses int
}

func parseInput(line string) ([]int, [][]int, []int) {
	state := make([]int, 0)
	buttons := make([][]int, 0)
	jolts := make([]int, 0)

	items := strings.SplitSeq(line, " ")
	for item := range items {
		switch item[0] {
		case '[': // state
			st := strings.Trim(item, "[]")
			for _, ch := range st {
				switch ch {
				case '.':
					state = append(state, 0)
				case '#':
					state = append(state, 1)
				}
			}
		case '(': // button
			bts := strings.Split(strings.Trim(item, "()"), ",")
			button := make([]int, 0)
			for _, bt := range bts {
				button = append(button, utils.Atoi(bt))
			}
			buttons = append(buttons, button)
		case '{': // jolt
			jts := strings.Split(strings.Trim(item, "{}"), ",")
			for _, jt := range jts {
				jolts = append(jolts, utils.Atoi(jt))
			}
		}

	}

	return state, buttons, jolts
}

func bfs(inital, target []int, buttons [][]int) int {
	stateToInt := func(state []int) int {
		res := 0
		for i, v := range state {
			if v == 1 {
				res |= (1 << i)
			}
		}
		return res
	}

	btnMasks := make([]int, len(buttons))
	for i, btn := range buttons {
		mask := 0
		for _, pos := range btn {
			mask |= (1 << pos)
		}
		btnMasks[i] = mask
	}

	initState := stateToInt(inital)
	targetState := stateToInt(target)

	if initState == targetState {
		return 0
	}

	visited := make(map[int]bool)
	visited[initState] = true
	queue := []QueueItem{{state: initState, presses: 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, btnMask := range btnMasks {
			newState := current.state ^ btnMask

			if newState == targetState {
				return current.presses + 1
			}

			if !visited[newState] {
				visited[newState] = true
				queue = append(queue, QueueItem{state: newState, presses: current.presses + 1})
			}
		}
	}

	return -1 //Impossible
}

func part1() {
	lines := make([]string, 0)
	utils.ReadFile("./cmd/day10/input.txt", func(line string) {
		lines = append(lines, line)
	})

	sum := 0
	for _, line := range lines {
		state, buttons, _ := parseInput(line)
		initState := make([]int, len(state))
		presses := bfs(initState, state, buttons)
		sum += presses
	}

	fmt.Println("Part 1:", sum)

}

func solveILP(coeff [][]int, diff []int, numButtons int, maxSinglePress int) int {
	n := len(diff)
	bestResult := math.MaxInt32

	var solve func(buttonIndex int, currentPress []int, totalPress int)
	solve = func(buttonIndex int, currentPress []int, totalPress int) {
		if totalPress >= bestResult {
			return
		}
		if buttonIndex == numButtons {
			for i := range n {
				sum := 0
				for j := range numButtons {
					sum += coeff[i][j] * currentPress[j]
				}
				if sum != diff[i] {
					return
				}
			}
			bestResult = totalPress
			return
		}

		maxForButton := maxSinglePress
		for i := range n {
			if coeff[i][buttonIndex] > 0 {
				currntSum := 0
				for j := range buttonIndex {
					currntSum += coeff[i][j] * currentPress[j]
				}
				remaining := diff[i] - currntSum
				if remaining < 0 {
					return
				}
				if remaining < maxForButton {
					maxForButton = remaining
				}
			}
		}

		for presses := 0; presses <= maxForButton; presses++ {
			currentPress[buttonIndex] = presses
			solve(buttonIndex+1, currentPress, totalPress+presses)
		}
	}

	currentPress := make([]int, numButtons)
	solve(0, currentPress, 0)

	if bestResult == math.MaxInt32 {
		return -1
	}
	return bestResult
}

func minButtonPresses(initial, target []int, buttons [][]int) int {
	n := len(initial)
	numButtons := len(buttons)

	diff := make([]int, n)
	for i := range n {
		diff[i] = target[i] - initial[i]
		if diff[i] < 0 {
			return -1
		}
	}

	allZero := true
	for _, d := range diff {
		if d != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}

	coeff := make([][]int, n)
	for i := range n {
		coeff[i] = make([]int, numButtons)
	}
	for j, btn := range buttons {
		for _, pos := range btn {
			if pos < n {
				coeff[pos][j] = 1
			}
		}
	}

	maxPresses := 0
	for _, d := range diff {
		if d > maxPresses {
			maxPresses = d
		}
	}

	result := solveILP(coeff, diff, numButtons, maxPresses)
	return result
}

func part2() {
	lines := make([]string, 0)
	utils.ReadFile("./cmd/day10/input.txt", func(line string) {
		lines = append(lines, line)
	})

	sum := 0
	for _, line := range lines {
		_, buttons, jolts := parseInput(line)
		initState := make([]int, len(jolts))
		presses := minButtonPresses(initState, jolts, buttons)
		sum += presses
	}

	fmt.Println("Part 2:", sum)
}

func main() {
	part1()
	part2()
}
