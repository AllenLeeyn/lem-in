package main

import "fmt"

func (m *maze) getMoving() {
	for exitCnt, turns := 0, 0; exitCnt < m.antQty; turns++ {
		fmt.Printf("Turns %v:", turns+1)
		for i, path := range m.sol.paths {
			if m.paths[path].antCnt == 0 {
				continue
			}
			lastIndex := len(m.paths[path].seq) - 1
			lastRoom := m.paths[path].seq[lastIndex]
			for _, room := range m.paths[path].seq {
				if m.rooms[room].antNm == 0 {
					fmt.Printf(" L%v-%s", i+1, room)
					m.rooms[room].antNm = i + 1
					break
				}
				if room == lastRoom {
					fmt.Printf(" L%v-%s", m.rooms[room].antNm, m.end)
					m.paths[path].antCnt--
					exitCnt++
				}
				m.rooms[room].antNm += len(m.sol.paths)
				if m.rooms[room].antNm != 0 && m.rooms[room].antNm < m.antQty {
					fmt.Printf(" L%v-%s", m.rooms[room].antNm, room)
				}
			}
		}
		fmt.Println()
	}
}

func (m *maze) getAntCnt() {
	antQty := m.antQty
	longestPath := m.paths[m.sol.paths[0]]
	for _, path := range m.sol.paths {
		m.paths[path].antCnt = longestPath.length - m.paths[path].length
		antQty -= (longestPath.length - m.paths[path].length)
	}
	pathCnt := len(m.sol.paths)
	for _, path := range m.sol.paths {
		m.paths[path].antCnt += antQty / pathCnt
		antQty -= antQty / pathCnt
		pathCnt--

	}
}
