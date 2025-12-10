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

var cache = make(map[Point]bool)

func inPolygon(pt Point, poly []Point) bool {

	if v, ok := cache[pt]; ok {
		return v
	}

	// if !pointInboundingBox(pt) {
	// 	cache[pt] = false
	// 	return false
	// }

	size := len(poly)
	x := pt.x
	y := pt.y
	inside := false

	p1 := poly[0]
	var p2 Point

	for i := 1; i <= size; i++ {
		p2 = poly[i%size]

		if (x == p1.x && x == p2.x && y >= utils.Min(p1.y, p2.y) && y <= utils.Max(p1.y, p2.y)) ||
			y == p1.y && y == p2.y && x >= utils.Min(p1.x, p2.x) && x <= utils.Max(p1.x, p2.x) {
			cache[pt] = true
			return true
		}

		if ((p2.y > y) != (p1.y > y)) &&
			(x < (y-p1.y)*(p2.x-p1.x)/(p2.y-p1.y)+p1.x) {
			inside = !inside
		}

		p1 = p2
	}

	cache[pt] = inside

	return inside
}

func edgeIntersectSquare(points []Point, a, b Point) bool {
	size := len(points)
	xs := []int{a.x, b.x}
	ys := []int{a.y, b.y}
	slices.Sort(xs)
	slices.Sort(ys)
	x1, x2, y1, y2 := xs[0], xs[1], ys[0], ys[1]

	p1 := points[0]
	var p2 Point

	for i := 1; i <= size; i++ {
		p2 = points[i%size]
		if y1 < p1.y && p1.y < y2 {
			if utils.Max(p1.x, p2.x) > x1 && utils.Min(p1.x, p2.x) < x2 {
				return true
			}
		} else {
			if x1 < p1.x && p1.x < x2 {
				if utils.Max(p1.y, p2.y) > y1 && utils.Min(p1.y, p2.y) < y2 {
					return true
				}
			}
		}

		p1 = p2
	}

	return false
}

func areaIsValid(p1, p2, p3, p4 Point, points []Point) bool {
	for _, p := range []Point{p1, p2, p3, p4} {
		if !inPolygon(p, points) {
			return false
		}
	}

	if edgeIntersectSquare(points, p1, p2) {
		return false
	}

	return true
}

var (
	bottom, top Point
)

func part2() {
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

			p3 := Point{p.x, p1.y}
			p4 := Point{p1.x, p.y}

			lx := p1.x - p.x + 1
			ly := utils.Abs(p1.y-p.y) + 1

			s := lx * ly
			if areaIsValid(p, p1, p3, p4, points) {
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
