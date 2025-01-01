package main

import "fmt"

func (m *maze) getMoving() {
	// for there are ants not out yet
	// counting turns too
	nextAnt, prevAnt := 1, 0
	for exitCnt, turns := 0, 0; exitCnt < m.antQty; turns++ {
		fmt.Printf("Turn %03d:", turns+1)

		// for each path in solution
		for _, path := range m.sol.paths {
			// if ants assigned are processed, path is done
			if m.paths[path].antsProcessed == m.paths[path].antsAssigned {
				continue
			}

			if len(m.paths[path].seq) == 0 {
				fmt.Printf(" L%v-%s", nextAnt, m.end)
				m.paths[path].antProcessing++
				m.paths[path].antsProcessed++
				exitCnt++
				nextAnt++
				continue
			}

			// get last room to know when to print exit
			lastIndex := len(m.paths[path].seq) - 1
			lastRoom := m.paths[path].seq[lastIndex]

			// for each room in path
			for i, room := range m.paths[path].seq {
				// if room is the last room
				if room == lastRoom && m.rooms[room].antNm != 0 {
					fmt.Printf(" L%v-%s", m.rooms[room].antNm, m.end)
					m.paths[path].antsProcessed++
					exitCnt++
				}

				if i == 0 && m.paths[path].antProcessing < m.paths[path].antsAssigned {
					prevAnt, m.rooms[room].antNm = m.rooms[room].antNm, nextAnt
					fmt.Printf(" L%v-%s", nextAnt, room)
					m.paths[path].antProcessing++
					nextAnt++
					continue
				} else if i == 0 {
					prevAnt, m.rooms[room].antNm = m.rooms[room].antNm, 0
					continue
				}

				m.rooms[room].antNm, prevAnt = prevAnt, m.rooms[room].antNm
				if m.rooms[room].antNm != 0 {
					fmt.Printf(" L%v-%s", m.rooms[room].antNm, room)
				}
			}
		}
		fmt.Println()
	}
}

func (m *maze) getAntsAssignment() {
	antQty := m.antQty
	longestPath := m.paths[m.sol.paths[0]]
	for _, path := range m.sol.paths {
		m.paths[path].antsAssigned = longestPath.length - m.paths[path].length
		antQty -= (longestPath.length - m.paths[path].length)
	}
	pathCnt := len(m.sol.paths)
	for _, path := range m.sol.paths {
		m.paths[path].antsAssigned += antQty / pathCnt
		antQty -= antQty / pathCnt
		pathCnt--

	}
}
