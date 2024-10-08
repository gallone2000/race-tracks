package main

import (
	"assignment/pkg/shortestpath"
	"fmt"
)

func main() {
	var numTestCases int
	fmt.Print("Enter the number of test cases: ")
	fmt.Scan(&numTestCases)

	for i := 0; i < numTestCases; i++ {
		var X, Y int
		for {
			fmt.Printf("Enter grid dimensions for test case %d (X Y): ", i+1)
			fmt.Scan(&X, &Y)
			if X >= 1 && X <= 30 && Y >= 1 && Y <= 30 {
				break
			} else {
				fmt.Println("Invalid dimensions. Please enter valid values (1 <= X <= 30 and 1 <= Y <= 30).")
			}
		}

		var startX, startY, finishX, finishY int
		fmt.Printf("Enter start and finish points for test case %d (startX startY finishX finishY): ", i+1)
		fmt.Scan(&startX, &startY, &finishX, &finishY)
		start := shortestpath.State{X: startX, Y: startY}
		finish := shortestpath.State{X: finishX, Y: finishY}

		var numObstacles int
		fmt.Printf("Enter the number of obstacles for test case %d: ", i+1)
		fmt.Scan(&numObstacles)

		var obstacles []shortestpath.RawObstacle
		for j := 0; j < numObstacles; j++ {
			var x1, y1, x2, y2 int
			for {
				fmt.Printf("Enter obstacle %d coordinates (x1 x2 y1 y2): ", j+1)
				_, err := fmt.Scan(&x1, &x2, &y1, &y2)
				if err != nil {
					fmt.Println("Invalid input. Please enter valid coordinates (x1 x2 y1 y2).")
					continue
				} else {
					obstacles = append(obstacles, shortestpath.RawObstacle{X1: x1, X2: x2, Y1: y1, Y2: y2})
					break
				}
			}
		}

		result, err := shortestpath.CalculateShortestPath(start, finish, X, Y, obstacles)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			if result == -1 {
				fmt.Println("No solution.")
			} else {
				fmt.Printf("Optimal solution for test case %d takes %d hops.\n", i+1, result)
			}
		}
	}
}
