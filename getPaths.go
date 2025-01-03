package main

// m.getPaths() get any path that connects the start to end.
func (m *maze) getPaths(curPath []string) {
	length := len(curPath)
	curRoom := curPath[length-1]
	for _, nextRoom := range m.rooms[curRoom].linkTo {
		if nextRoom == m.end {
			m.paths = append(m.paths,
				pathStruct{seq: curPath[1:], length: length})
			continue
		}
		if !isVisited(curPath, nextRoom) {
			m.getPaths(append(curPath, nextRoom))
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
