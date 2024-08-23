package shortestpath

import (
	"errors"
	"fmt"
)

var (
	minVelocity = -3
	maxVelocity = 3
)

type State struct {
	X, Y, Vx, Vy int
}

var accelerations = []struct{ dx, dy int }{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 0}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func CalculateShortestPath(start, finish State, X, Y int, obstacles map[string]bool) (int, error) {
	if !gridSizeIsValid(X, Y) {
		return -1, errors.New("invalid grid size")
	}
	if !ObstaclesAreValid(obstacles, start, finish) {
		return -1, errors.New("invalid obstacles. Start and finish points cannot be obstacles")
	}

	queue := []State{start}
	visited := make(map[State]bool)
	minHops := make(map[State]int)
	parents := make(map[State]State)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.X == finish.X && curr.Y == finish.Y {
			return minHops[curr], nil
		}

		for _, nextState := range generateNextStates(curr, X, Y) {
			if !visited[nextState] && !obstacles[fmt.Sprintf("%d.%d", nextState.X, nextState.Y)] {
				queue = append(queue, nextState)
				visited[nextState] = true
				parents[nextState] = curr
				minHops[nextState] = minHops[curr] + 1
			}
		}
	}

	return -1, nil // No solution
}

func isValid(x, y, X, Y int) bool {
	return x >= 0 && x < X && y >= 0 && y < Y
}

func velocityIsValid(vx, vy int) bool {
	return (minVelocity <= vx && vx <= maxVelocity) && (minVelocity <= vy && vy <= maxVelocity)
}

func generateNextStates(curr State, X, Y int) []State {
	var nextStates []State
	for _, acc := range accelerations {
		vx, vy := curr.Vx+acc.dx, curr.Vy+acc.dy
		if !velocityIsValid(vx, vy) {
			continue
		}
		x, y := curr.X+vx, curr.Y+vy

		if isValid(x, y, X, Y) {
			nextStates = append(nextStates, State{x, y, vx, vy})
		}
	}

	return nextStates
}
