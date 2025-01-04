package main

import "fmt"

// m.getMoving() moves ants along each path
func (m *maze) getMoving() {
	nextAntID, prevAnt := 1, 0

	for exitQty, turns := 0, 0; exitQty < m.antQty; turns++ {
		m.movement += fmt.Sprintf("Turn %03d:", turns+1)
		for _, pID := range m.sol.pathIDs {
			if m.paths[pID].antsProcessed == m.paths[pID].antsAssigned {
				continue
			}
			for i, room := range m.paths[pID].seq {
				if i == 0 && m.paths[pID].antProcessing < m.paths[pID].antsAssigned {
					m.rooms[room].antID, prevAnt = nextAntID, m.rooms[room].antID
					m.paths[pID].antProcessing++
					nextAntID++
				} else if i == 0 {
					m.rooms[room].antID, prevAnt = 0, m.rooms[room].antID
				} else {
					m.rooms[room].antID, prevAnt = prevAnt, m.rooms[room].antID
				}
				exitQty = m.getMove(room, m.rooms[room].antID, pID, exitQty)
			}
		}
		m.movement += "\n"
	}
}

// m.getMove() c
func (m *maze) getMove(room string, aID, pID, exitQty int) int {
	if aID != 0 {
		color := colors[aID%12]
		if room == m.end {
			color = "\033[7m" + color
			m.paths[pID].antsProcessed++
			exitQty++
		}
		m.movement += fmt.Sprintf(" %sL%v-%s%s", color, aID, room, resetColor)
	}
	return exitQty
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
