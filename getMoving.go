package main

import "fmt"

// m.getMoving() moves ants along each path
func (m *maze) getMoving() {
	nextAnt, prevAnt := 1, 0
	for exitQty, turns := 0, 0; exitQty < m.antQty; turns++ {
		fmt.Printf("Turn %03d:", turns+1)

		for _, id := range m.sol.pathIDs {
			if m.paths[id].antsProcessed == m.paths[id].antsAssigned {
				continue
			}
			if len(m.paths[id].seq) == 0 {
				exitQty = m.antsExiting(&m.paths[id], nextAnt, exitQty)
				m.paths[id].antProcessing++
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
					exitQty = m.antsExiting(&m.paths[id], m.rooms[room].antID, exitQty)
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
func (m *maze) antsExiting(path *pathStruct, antID, exitQty int) int {
	fmt.Printf(" L%v-%s", antID, m.end)
	path.antsProcessed++
	return exitQty + 1
}

// m.getAntAssignment() assigns ants to each path base on path's length
func (m *maze) getAntsAssignment() {
	antQty := m.antQty
	longestPath := m.paths[m.sol.pathIDs[0]]
	for _, id := range m.sol.pathIDs {
		m.paths[id].antsAssigned = longestPath.length - m.paths[id].length
		antQty -= (longestPath.length - m.paths[id].length)
	}
	pathCnt := len(m.sol.pathIDs)
	for _, id := range m.sol.pathIDs {
		m.paths[id].antsAssigned += antQty / pathCnt
		antQty -= antQty / pathCnt
		pathCnt--

	}
}
