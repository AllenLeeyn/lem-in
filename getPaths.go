package main

import "sort"

// m.getPaths() get any path that connects the start to end.
func (m *maze) getPaths(curPath []string) {
	length := len(curPath)
	curRoom := curPath[length-1]
	newPath := append([]string{}, curPath...)
	for _, nextRoom := range m.rooms[curRoom].linkTo { //loop through the links from the current room
		if nextRoom == m.end {  //current path is finshed
			m.paths = append(m.paths,
				pathStruct{seq: newPath[1:], length: length})
			continue
		}
		if !isVisited(newPath, nextRoom) {
			m.getPaths(append(newPath, nextRoom))
		}
	}
}

// isVisited() checkes if the nextRoom is visited in the curPath
func isVisited(curPath []string, nextRoom string) bool {
	for _, visited := range curPath {
		if visited == nextRoom {
			return true
		}
	}
	return false
}

// m.sortPaths() sorts m.paths in ascending order based on
func (m *maze) sortPaths() {
	sort.Slice(m.paths, func(i, j int) bool {
		return m.paths[i].length < m.paths[j].length
	})
}
