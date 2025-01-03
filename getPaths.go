package main

// m.getPaths() get any path that connects the start to end.
func (m *maze) getPaths(curPath []string) {
	length := len(curPath)
	curRoom := curPath[length-1]
	newPath := append([]string{}, curPath...)
	for _, nextRoom := range m.rooms[curRoom].linkTo {
		if nextRoom == m.end {
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
