package main

import (
	"fmt"
	"slices"
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

func checkStatesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func bfs2(initState, target []int, buttons [][]int) int {
	stateToStr := func(state []int) string {
		res := ""
		for _, v := range state {
			res += fmt.Sprintf("%d,", v)
		}
		return res
	}

	visited := make(map[string]bool)
	visited[stateToStr(initState)] = true
	queue := []QueueItem2{{state: initState, presses: 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, btn := range buttons {
			newState := slices.Clone(current.state)

			for _, pos := range btn {
				newState[pos] += 1
			}

			if checkStatesEqual(newState, target) {
				return current.presses + 1
			}

			stateInt := stateToStr(newState)
			if !visited[stateInt] {
				visited[stateInt] = true
				queue = append(queue, QueueItem2{state: newState, presses: current.presses + 1})
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

func part2() {
	lines := make([]string, 0)
	utils.ReadFile("./cmd/day10/input.txt", func(line string) {
		lines = append(lines, line)
	})

	sum := 0
	for _, line := range lines {
		_, buttons, jolts := parseInput(line)
		initState := make([]int, len(jolts))
		presses := bfs2(initState, jolts, buttons)
		sum += presses
		break
	}

	fmt.Println("Part 2:", sum)
}

func main() {
	part1()
	part2()
}
