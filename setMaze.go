package main

import (
	"fmt"
	"strconv"
	"strings"
)

// m.printMaze() prints the input given
func (m *maze) printMaze() {
	for _, line := range m.result {
		fmt.Println(line)
	}
	fmt.Println()
}

// m.setMaze() process the fileInput set the maze definition/parameters
func (m *maze) setMaze(fileInput []string) error {
	if err := m.setAntQty(fileInput[0]); err != nil {
		return err
	}
	if err := m.setRooms(fileInput); err != nil {
		return err
	}
	if err := m.setLinks(fileInput); err != nil {
		return err
	}
	return nil
}

// m.setAntQty() process the first line of fileInput for antQty
func (m *maze) setAntQty(fileInput string) error {
	antQty, err := strconv.Atoi(fileInput)
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format; %s is not a valid ant quantity",
			fileInput)
	} else if antQty < 1 { //check too few ants
		return fmt.Errorf("ERROR: %s is not a valid ant quantity",
			fileInput)
	}
	m.antQty, m.result = antQty, append(m.result, fileInput)
	return nil
}

// m.setRooms() process the whole inputFile for valid room definition.
// it records a room as the start/end base if a proper declaration is
// given in the previous line (##start/ ##end).
func (m *maze) setRooms(fileInput []string) error {
	startCount, endCount, found := 0, 0, ""

	for i := 0; i < len(fileInput); i++ {
		if fileInput[i] == "##start" || fileInput[i] == "##end" {
			m.result = append(m.result, fileInput[i])
			found = fileInput[i]
			if i+1 < len(fileInput) {
				i++
			}
		}

		if !strings.Contains(fileInput[i], " ") {
			continue
		}
		name, isOk := m.setRoom(fileInput[i])
		if !isOk {
			return fmt.Errorf("ERROR: invalid data format. Check room values: %s", fileInput[i])
		}
		switch found {
		case "##start":
			m.start, found = name, ""
			startCount++
		case "##end":
			m.end, found = name, ""
			endCount++
		}
		m.result = append(m.result, fileInput[i])
	}
	if startCount != 1 || endCount != 1 {
		return fmt.Errorf("ERROR: invalid data format. Check quantity of start/end rooms")
	}
	return nil
}

// m.setRoom() process a room defintion line.
// It accepts 3 inputs split by space, and no '-' is allowed in the line.
// It also checks if there is repeat of room name or coordinates.
func (m *maze) setRoom(rmValues string) (string, bool) {
	values := strings.Fields(rmValues)
	if len(values) != 3 || strings.Contains(rmValues, "-") {
		return "", false
	}
	name := values[0]
	if name[0] == 'L' || name[0] == '#' {
		return "", false
	}
	x, err := strconv.Atoi(values[1])
	if err != nil {
		return "", false
	}
	y, err := strconv.Atoi(values[2])
	if err != nil {
		return "", false
	}
	for key, room := range m.rooms {
		if name == key || (room.x == x && room.y == y) {
			return "", false
		}
	}
	m.rooms[name] = &room{x: x, y: y}
	return name, true
}

// m.setLinks() process the whole inputFile for valid link definition.s
func (m *maze) setLinks(fileInput []string) error {
	for i := 0; i < len(fileInput); i++ {
		if !strings.Contains(fileInput[i], "-") {
			continue
		}
		link := strings.Split(fileInput[i], "-")
		if len(link) != 2 || !m.isValidLink(link[0], link[1]) {
			return fmt.Errorf("ERROR: invalid data format. Invalid link: %s", fileInput[i])
		}
		m.rooms[link[0]].linkTo = append(m.rooms[link[0]].linkTo, link[1])
		m.rooms[link[1]].linkTo = append(m.rooms[link[1]].linkTo, link[0])
		m.result = append(m.result, fileInput[i])
	}
	return nil
}

// m.isValidLink() checks if link given is valid.
// It checks if rooms in link exist and if the link is already establish.
func (m *maze) isValidLink(roomA, roomB string) bool {
	_, isExistA := m.rooms[roomA]
	room, isExistB := m.rooms[roomB]
	if !isExistA || !isExistB || roomA == roomB {
		return false
	}
	for _, existed := range room.linkTo {
		if roomA == existed {
			return false
		}
	}
	return true
}
