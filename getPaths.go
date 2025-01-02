package main

func (m *maze) getPaths(curPath []string) {
	length := len(curPath)
	curRoom := curPath[length-1]
	if curRoom == m.end {
		m.paths = append(m.paths,
			pathStruct{seq: curPath[1 : length-1], length: length - 1})
		return
	}
	newPath := append([]string{}, curPath...)
	for _, nextRoom := range m.rooms[curRoom].linkTo {
		if !isVisited(curPath, nextRoom) {
			m.getPaths(append(newPath, nextRoom))
		}
	}
}

func isVisited(curPath []string, nextRoom string) bool {
	for _, visited := range curPath {
		if visited == nextRoom {
			return true
		}
	}
	return false
}
