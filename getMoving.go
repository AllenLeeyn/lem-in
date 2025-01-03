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
			for i, room := range m.paths[id].seq {
				if i == 0 && m.paths[id].antProcessing < m.paths[id].antsAssigned {
					prevAnt, m.rooms[room].antID = m.rooms[room].antID, nextAnt
					color := colors[m.rooms[room].antID%12]
					fmt.Printf(color+" L%v-%s"+resetColor, nextAnt, room)
					m.paths[id].antProcessing++
					nextAnt++
				} else if i == 0 {
					prevAnt, m.rooms[room].antID = m.rooms[room].antID, 0
				}
				if i == 0 && room == m.end {
					m.paths[id].antsProcessed++
					exitQty++
					continue
				}

				if i != 0 && room != m.end {
					m.rooms[room].antID, prevAnt = prevAnt, m.rooms[room].antID
					if m.rooms[room].antID != 0 {
						color := colors[m.rooms[room].antID%12]
						fmt.Printf(color+" L%v-%s"+resetColor, m.rooms[room].antID, room)
					}
				}

				if room == m.end && prevAnt != 0 {
					color := colors[prevAnt%12]
					fmt.Printf(color+" L%v-%s"+resetColor, prevAnt, m.end)
					m.paths[id].antsProcessed++
					exitQty++
				}
			}
		}
		fmt.Println()
	}
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
