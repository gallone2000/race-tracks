package shortestpath

import (
	"testing"
)

func TestBfsWithShortestPath(t *testing.T) {
	X, Y := 5, 5
	start := State{X: 4, Y: 0}
	finish := State{X: 4, Y: 4}
	obstacles := make(map[string]bool)
	GenerateObstacles(&obstacles, 1, 2, 4, 3)
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
	obstacles := make(map[string]bool)
	GenerateObstacles(&obstacles, 1, 2, 4, 3)
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
	obstacles := make(map[string]bool)
	GenerateObstacles(&obstacles, 1, 0, 1, 2)
	GenerateObstacles(&obstacles, 0, 1, 2, 1)
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
	obstacles := make(map[string]bool)
	GenerateObstacles(&obstacles, 0, 0, 1, 2)
	GenerateObstacles(&obstacles, 0, 1, 2, 1)
	result, err := CalculateShortestPath(start, finish, X, Y, obstacles)
	expected := -1

	if err == nil {
		t.Error("Expect error")
	}
	if result != expected {
		t.Errorf("Expected %d hops, but got %d", expected, result)
	}
	if err.Error() != "invalid obstacles. Start and finish points cannot be obstacles" {
		t.Error("Expected error message")
	}
}
