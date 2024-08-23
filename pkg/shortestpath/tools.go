package shortestpath

import "fmt"

var (
	maxXgridSize = 30
	maxYgridSize = 30
)

func GenerateObstacles(obstacles *map[string]bool, x1, y1, x2, y2 int) *map[string]bool {
	actualObstacles := *obstacles

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			actualObstacles[fmt.Sprintf("%d.%d", x, y)] = true
		}
	}

	return &actualObstacles
}

func ObstaclesAreValid(obstacles map[string]bool, start, finish State) bool {
	obstaclesAreValid := true
	startCoordinates := fmt.Sprintf("%d.%d", start.X, start.Y)
	endCoordinates := fmt.Sprintf("%d.%d", finish.X, finish.Y)
	for obstacleXY, _ := range obstacles {
		if obstacleXY == startCoordinates || obstacleXY == endCoordinates {
			return false
		}
	}
	return obstaclesAreValid
}

func gridSizeIsValid(X, Y int) bool {
	return X <= maxXgridSize && Y <= maxYgridSize
}
