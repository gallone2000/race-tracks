package shortestpath

import (
	"errors"
	"fmt"
)

var (
	maxXgridSize = 30
	maxYgridSize = 30
)

func generateObstacles(obstacles *map[string]bool, obstacle RawObstacle, X, Y int, start, finish State) error {
	actualObstacles := *obstacles
	if !obstacleCoordinatesAreValid(X, Y, start, finish, obstacle) {
		return errors.New("obstacle coordinates are invalid")
	}
	for x := obstacle.X1; x <= obstacle.X2; x++ {
		for y := obstacle.Y1; y <= obstacle.Y2; y++ {
			actualObstacles[fmt.Sprintf("%d.%d", x, y)] = true
		}
	}
	obstacles = &actualObstacles

	return nil
}

func gridSizeIsValid(X, Y int) bool {
	return X <= maxXgridSize && Y <= maxYgridSize
}

func obstacleCoordinatesAreValid(X, Y int, start, finish State, obstacle RawObstacle) bool {
	o1 := State{X: obstacle.X1, Y: obstacle.Y1}
	o2 := State{X: obstacle.X2, Y: obstacle.Y2}

	return obstacle.X1 <= obstacle.X2 &&
		0 <= obstacle.X1 &&
		obstacle.X2 < X &&
		obstacle.Y1 <= obstacle.Y2 &&
		0 <= obstacle.Y1 &&
		obstacle.Y2 < Y &&
		o1 != start &&
		o2 != finish
}

func startFinishAreValid(X, Y int, start, finish State) bool {
	return 0 <= start.X &&
		0 <= finish.X &&
		start.X < X &&
		finish.X < X &&
		0 <= start.Y &&
		0 <= finish.Y &&
		start.Y < Y &&
		finish.Y < Y
}
