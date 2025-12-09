package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/umbe77/aoc-2025/utils"
)

type Point struct {
	x, y int
}

func part1() {
	points := make([]Point, 0)
	utils.ReadFile("./cmd/day09/input.txt", func(line string) {
		p := strings.Split(line, ",")
		points = append(points, Point{utils.Atoi(p[0]), utils.Atoi(p[1])})
	})

	square := 0

	for _, p := range points {
		for _, p1 := range points {
			if p == p1 {
				continue
			}

			lx := p1.x - p.x + 1
			ly := utils.Abs(p1.y-p.y) + 1

			s := lx * ly
			// fmt.Printf("p: %v, p1: %v, s: %d\n", p, p1, s)
			if s > square {
				square = s
			}

		}
	}

	fmt.Println("Part 1", square)
}

func inPolygon(pt Point, poly []Point) bool {

	if slices.Index(poly, pt) != -1 {
		return true
	}

	size := len(poly)
	x := pt.x
	y := pt.y
	inside := false

	p1 := poly[0]
	var p2 Point

	for i := 1; i <= size; i++ {
		p2 = poly[i%size]

		if y > utils.Min(p1.y, p2.y) {
			if y <= utils.Max(p1.y, p2.y) {
				if x <= utils.Max(p1.x, p2.x) {
					intersect := (y-p1.y)*(p2.x-p1.x)/(p2.y-p1.y) + p1.x
					if p1.x == p2.x || x <= intersect {
						inside = !inside
					}
				}
			}
		}

		p1 = p2
	}

	return inside
}

func part2() {
	points := make([]Point, 0)
	utils.ReadFile("./cmd/day09/input.txt", func(line string) {
		p := strings.Split(line, ",")
		points = append(points, Point{utils.Atoi(p[0]), utils.Atoi(p[1])})
	})

	square := 0

	for _, p := range points {
		for _, p1 := range points {
			if p == p1 || p.y == p1.y || p.x == p1.x {
				continue
			}

			p3 := Point{p.x, p1.y}
			p4 := Point{p1.x, p.y}

			if inPolygon(p3, points) && inPolygon(p4, points) {
				lx := p1.x - p.x + 1
				ly := utils.Abs(p1.y-p.y) + 1

				s := lx * ly
				if s > square {
					square = s
				}
			}

		}
	}

	fmt.Println("Part 2", square)
}

func main() {
	part1()
	part2()
}
