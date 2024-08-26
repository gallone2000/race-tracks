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

type RawObstacle struct {
	X1, X2, Y1, Y2 int
}

var accelerations = []struct{ dx, dy int }{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 0}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func CalculateShortestPath(start, finish State, X, Y int, rawObstacles []RawObstacle) (int, error) {
	if !gridSizeIsValid(X, Y) {
		return -1, errors.New("invalid grid size")
	}
	if !startFinishAreValid(X, Y, start, finish) {
		return -1, errors.New("start and finish coordinates are invalid")
	}
	obstacles := make(map[string]bool)
	for _, rawObstacle := range rawObstacles {
		err := generateObstacles(&obstacles, rawObstacle, X, Y, start, finish)
		if err != nil {
			return -1, err
		}
	}

	queue := []State{start}
	visited := make(map[State]bool)
	minHops := make(map[State]int)

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
