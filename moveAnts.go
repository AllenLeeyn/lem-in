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
	// Create an array of AntMovement structs to track each ant
	antMovements := make([]AntMovement, l.antQty)

	// Initialize ants with paths, each ant gets a path in a round-robin fashion, but more intelligently
	// Ants will be distributed across paths by considering their length (fewer ants in longer paths).
	antCountPerPath := make([]int, len(l.paths))

	// Distribute ants across the paths intelligently
	for i := 0; i < l.antQty; i++ {
		shortestPathIndex := getShortestPathIndex(antCountPerPath, l.paths) // Get the path with the least number of ants
		antMovements[i].antID = i + 1
		antMovements[i].path = l.paths[shortestPathIndex] // Assign ant to the shortest path
		antMovements[i].position = 1                      // start printing the room after the start room
		antCountPerPath[shortestPathIndex]++
	}

	// Start the simulation and print movements
	// Track the movement of ants in each step, ensuring no room is occupied twice
	for {
		allAntsFinished := true
		var movementOutput []string
		occupiedRooms := make(map[string]bool) // Map to track occupied rooms in each step

		// Check if any ants are still moving, and move them
		for i := 0; i < len(antMovements); i++ {
			ant := &antMovements[i]
			if ant.position < len(ant.path) {
				allAntsFinished = false
				room := ant.path[ant.position]

				// Check if the room is occupied by another ant
				if _, occupied := occupiedRooms[room]; !occupied {
					// Move the ant to the next room in the path
					movementOutput = append(movementOutput, fmt.Sprintf("%d-%s", ant.antID, room))
					occupiedRooms[room] = true // Mark this room as occupied
					ant.position++             // Move to the next room in the path
				}
			}
		}

		// Print the current movement of all ants in the format "1-a 2-b 3-c ..."
		if len(movementOutput) > 0 {
			fmt.Println(strings.Join(movementOutput, " "))
		}

		// If all ants have finished, stop the simulation
		if allAntsFinished {
			break
		}
	}
}

// Helper function to get the index of the path with the fewest ants
func getShortestPathIndex(antCountPerPath []int, paths [][]string) int {
	minAnts := int(^uint(0) >> 1) // Initialize to a large number (max int)
	minPathIndex := 0

	// Find the path with the least number of ants assigned
	for i, path := range paths {
		if len(path) < len(paths[minPathIndex]) && antCountPerPath[i] < minAnts {
			minAnts = antCountPerPath[i]
			minPathIndex = i
		}
	}

	return minPathIndex
}
