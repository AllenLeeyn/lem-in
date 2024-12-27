package main

import (
	"fmt"
	"strings"
)

type AntMovement struct {
	antID    int
	path     []string
	position int // tracks the current room index in the path
}

func (l *Lemz) SimulateAntMovement() {
	antMovements := make([]AntMovement, l.antQty)
	antCountPerPath := make([]int, len(l.paths))

	// Distribute ants across paths optimally
	for i := 0; i < l.antQty; i++ {
		optimalPathIndex := getOptimalPathIndex(antCountPerPath, l.paths)
		antMovements[i].antID = i + 1
		antMovements[i].path = l.paths[optimalPathIndex]
		antMovements[i].position = 1
		antCountPerPath[optimalPathIndex]++
	}

	// Simulate movement
	for {
		allAntsFinished := true
		var movementOutput []string
		occupiedRooms := make(map[string]bool) // Track occupied rooms in the current step

		for i := 0; i < len(antMovements); i++ {
			ant := &antMovements[i]
			if ant.position < len(ant.path) {
				allAntsFinished = false
				room := ant.path[ant.position]

				// Allow multiple ants in `l.start` and `l.end`
				if room == l.start || room == l.end || !occupiedRooms[room] {
					// Move the ant to the next room
					movementOutput = append(movementOutput, fmt.Sprintf("%d-%s", ant.antID, room))
					if room != l.start && room != l.end {
						occupiedRooms[room] = true // Mark room as occupied
					}
					ant.position++
				}
			}
		}

		// Print the movements for this step
		if len(movementOutput) > 0 {
			fmt.Println(strings.Join(movementOutput, " "))
		}

		// Stop simulation if all ants have finished
		if allAntsFinished {
			break
		}
	}

}

func getOptimalPathIndex(antCountPerPath []int, paths [][]string) int {
	minCost := int(^uint(0) >> 1) // Large number (max int)
	optimalPathIndex := 0

	for i, path := range paths {
		cost := (len(path) - 1) + antCountPerPath[i] // Cost = path length + current congestion
		if cost < minCost {
			minCost = cost
			optimalPathIndex = i
		}
	}

	return optimalPathIndex
}
