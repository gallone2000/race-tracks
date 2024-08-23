package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func isValid(x, y, X, Y int) bool {
	return x >= 0 && x < X && y >= 0 && y < Y
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func solve(grid [][]bool, start, finish Point) string {
	X, Y := len(grid), len(grid[0])
	dirs := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // Possible hops

	// Initialize minimum hops array
	minHops := make([][]int, X)
	for i := range minHops {
		minHops[i] = make([]int, Y)
		for j := range minHops[i] {
			minHops[i][j] = math.MaxInt32
		}
	}

	queue := []Point{start}
	minHops[start.x][start.y] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			next := Point{curr.x + dir.x, curr.y + dir.y}
			if isValid(next.x, next.y, X, Y) && !grid[next.x][next.y] {
				newHops := minHops[curr.x][curr.y] + 1
				if newHops < minHops[next.x][next.y] {
					minHops[next.x][next.y] = newHops
					queue = append(queue, next)
				}
			}
		}
	}

	if minHops[finish.x][finish.y] == math.MaxInt32 {
		return "No solution."
	}
	return fmt.Sprintf("Optimal solution takes %d hops.", minHops[finish.x][finish.y])
}

func main() {
	// Example usage
	grid := make([][]bool, 5)
	for i := range grid {
		grid[i] = make([]bool, 5)
	}
	start := Point{4, 0}
	finish := Point{4, 4}
	grid[1][1] = true // Add the obstacle
	fmt.Println(solve(grid, start, finish))
}
