package main

import "fmt"

func (m *maze) getMoving() {
	// for there are ants not out yet
	// counting turns too
	nextAnt, prevAnt := 1, 0
	for exitCnt, turns := 0, 0; exitCnt < m.antQty; turns++ {
		fmt.Printf("Turn %03d:", turns+1)

		// for each path in solution
		for _, id := range m.sol.pathID {
			// if ants assigned are processed, path is done
			if m.paths[id].antsProcessed == m.paths[id].antsAssigned {
				continue
			}

			if len(m.paths[id].seq) == 0 {
				fmt.Printf(" L%v-%s", nextAnt, m.end)
				m.paths[id].antProcessing++
				m.paths[id].antsProcessed++
				exitCnt++
				nextAnt++
				continue
			}

			// get last room to know when to print exit
			lastIndex := len(m.paths[id].seq) - 1
			lastRoom := m.paths[id].seq[lastIndex]

			// for each room in path
			for i, room := range m.paths[id].seq {
				// if room is the last room
				if room == lastRoom && m.rooms[room].antID != 0 {
					fmt.Printf(" L%v-%s", m.rooms[room].antID, m.end)
					m.paths[id].antsProcessed++
					exitCnt++
				}

				if i == 0 && m.paths[id].antProcessing < m.paths[id].antsAssigned {
					prevAnt, m.rooms[room].antID = m.rooms[room].antID, nextAnt
					fmt.Printf(" L%v-%s", nextAnt, room)
					m.paths[id].antProcessing++
					nextAnt++
					continue
				} else if i == 0 {
					prevAnt, m.rooms[room].antID = m.rooms[room].antID, 0
					continue
				}

				m.rooms[room].antID, prevAnt = prevAnt, m.rooms[room].antID
				if m.rooms[room].antID != 0 {
					fmt.Printf(" L%v-%s", m.rooms[room].antID, room)
				}
			}
		}
		fmt.Println()
	}
}

func (m *maze) getAntsAssignment() {
	antQty := m.antQty
	longestPath := m.paths[m.sol.pathID[0]]
	for _, id := range m.sol.pathID {
		m.paths[id].antsAssigned = longestPath.length - m.paths[id].length
		antQty -= (longestPath.length - m.paths[id].length)
	}
	pathCnt := len(m.sol.pathID)
	for _, id := range m.sol.pathID {
		m.paths[id].antsAssigned += antQty / pathCnt
		antQty -= antQty / pathCnt
		pathCnt--

	}
}
