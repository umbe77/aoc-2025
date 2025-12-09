package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/umbe77/aoc-2025/utils"
)

type Point struct {
	x int
	y int
	z int
}

type Junction struct {
	d float64
	a Point
	b Point
}

func distance(a, b Point) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z

	s := dx*dx + dy*dy + dz*dz

	return math.Sqrt(float64(s))
}

func getHash(points []Point) string {
	slices.SortFunc(points, func(a, b Point) int {
		if a.x < b.x {
			return -1
		}
		if a.x > b.x {
			return 1
		}
		return 0
	})

	hash := fmt.Sprintf("%d,%d,%d||%d,%d,%d", points[0].x, points[0].y, points[0].z, points[1].x, points[1].y, points[1].z)
	return hash
}

func containsInJunction(junctions []Junction, a, b Point) bool {
	for _, v := range junctions {
		if v.a == a && v.b == b {
			return true
		}
		if v.b == a && v.a == b {
			return true
		}
	}
	return false
}

func part1() {
	boxes := make([]Point, 0)
	utils.ReadFile("./cmd/day08/sample.txt", func(line string) {
		coords := strings.Split(line, ",")
		boxes = append(boxes, Point{
			x: utils.Atoi(coords[0]),
			y: utils.Atoi(coords[1]),
			z: utils.Atoi(coords[2]),
		})
	})

	junctions := make([]Junction, 0)
	cache := make(map[string]float64)

	for i, p := range boxes {
		ds := 0.0
		for j, pc := range boxes {
			if i == j {
				continue
			}
			h := getHash([]Point{p, pc})
			d, ok := cache[h]
			if !ok {
				d = distance(p, pc)
				cache[h] = d
			}

			ds = d

			junction := Junction{
				d: ds,
				a: p,
				b: pc,
			}

			if !containsInJunction(junctions, p, pc) {
				junctions = append(junctions, junction)
			}

			// idx := slices.IndexFunc(junctions, func(j Junction) bool {
			// 	return (j.a == p || j.b == pc) && (j.a == pc || j.b == pc)
			// })
			// if idx == -1 {
			// 	junctions = append(junctions, junction)
			// }
		}
	}
	slices.SortFunc(junctions, func(a, b Junction) int {
		if a.d < b.d {
			return -1
		}
		if a.d > b.d {
			return 1
		}
		return 0
	})

	circuits := make([][]Point, 0)

	for i, v := range junctions {
		fmt.Println("junction", v)
		if i >= 11 {
			break
		}
		foundA := false
		foundB := false
		circuitIdx := -1
		for i, c := range circuits {
			idxa := slices.Index(c, v.a)
			idxb := slices.Index(c, v.b)

			if idxa == -1 && idxb == -1 {
				continue
			}

			if idxa == -1 && idxb > -1 {
				foundB = true
				circuitIdx = i
				break
			}

			if idxb == -1 && idxa > -1 {
				foundA = true
				circuitIdx = i
				break
			}
		}

		if circuitIdx != -1 {
			if !foundA {
				circuits[circuitIdx] = append(circuits[circuitIdx], v.a)
			}

			if !foundB {
				circuits[circuitIdx] = append(circuits[circuitIdx], v.b)
			}
		} else {
			circuits = append(circuits, []Point{v.a, v.b})
		}
		for _, c := range circuits {
			// if i >= 3 {
			// 	break
			// }
			fmt.Printf("%+v\n", c)
		}
		fmt.Println("======================")
	}
	slices.SortFunc(circuits, func(a, b []Point) int {
		la := len(a)
		lb := len(b)

		if la < lb {
			return 1
		}
		if la > lb {
			return -1
		}
		return 0
	})
	res := 1
	for i, c := range circuits {
		if i >= 3 {
			break
		}
		fmt.Printf("%+v\n", c)
		res *= len(c)
	}

	fmt.Println("Part 1", res)

}

func part2() {
	utils.ReadFile("./cmd/day08/sample.txt", func(line string) {
	})
}

func main() {
	part1()
	part2()
}
