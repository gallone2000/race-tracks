package shortestpath

import (
	"testing"
)

func TestBfsWithShortestPath(t *testing.T) {
	X, Y := 5, 5
	start := State{X: 4, Y: 0}
	finish := State{X: 4, Y: 4}
	var obstacles []RawObstacle
	obstacles = append(obstacles, RawObstacle{1, 4, 2, 3})
	result, err := CalculateShortestPath(start, finish, X, Y, obstacles)
	expected := 7

	if err != nil {
		t.Error("Expected no error, but got:", err)
	}
	if result != expected {
		t.Errorf("Expected %d hops, but got %d", expected, result)
	}
}

func TestBfsWithGridSizeError(t *testing.T) {
	X, Y := 31, 31
	start := State{X: 4, Y: 0}
	finish := State{X: 4, Y: 4}
	var obstacles []RawObstacle
	obstacles = append(obstacles, RawObstacle{1, 2, 4, 3})
	_, err := CalculateShortestPath(start, finish, X, Y, obstacles)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	expected := "invalid grid size"
	if err.Error() != expected {
		t.Errorf("Expected %s error, but got %s", expected, err.Error())
	}
}

func TestBfsWithNoShortestPath(t *testing.T) {
	X, Y := 3, 3
	start := State{X: 0, Y: 0}
	finish := State{X: 2, Y: 2}
	var obstacles []RawObstacle
	obstacles = append(obstacles, RawObstacle{1, 1, 0, 2})
	obstacles = append(obstacles, RawObstacle{0, 2, 1, 1})
	result, err := CalculateShortestPath(start, finish, X, Y, obstacles)
	expected := -1

	if err != nil {
		t.Error("Expected no error, but got:", err)
	}
	if result != expected {
		t.Errorf("Expected %d hops, but got %d", expected, result)
	}
}

func TestBfsWithError(t *testing.T) {
	X, Y := 3, 3
	start := State{X: 0, Y: 0}
	finish := State{X: 2, Y: 2}
	var obstacles []RawObstacle
	obstacles = append(obstacles, RawObstacle{0, 0, 1, 1})
	obstacles = append(obstacles, RawObstacle{1, 2, 2, 2})
	result, err := CalculateShortestPath(start, finish, X, Y, obstacles)
	expected := -1

	if err == nil {
		t.Error("Expect error")
	}
	if result != expected {
		t.Errorf("Expected %d hops, but got %d", expected, result)
	}
	if err.Error() != "obstacle coordinates are invalid" {
		t.Error("Expected error message")
	}
}

func TestErrorStartFinish(t *testing.T) {
	X, Y := 5, 5
	start := State{X: 5, Y: 0}
	finish := State{X: 4, Y: 6}
	var obstacles []RawObstacle
	obstacles = append(obstacles, RawObstacle{1, 4, 2, 3})
	_, err := CalculateShortestPath(start, finish, X, Y, obstacles)

	expected := "start and finish coordinates are invalid"
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err.Error() != expected {
		t.Errorf("Expected %s error, but got %s", expected, err.Error())
	}
}
