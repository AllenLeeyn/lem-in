package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (m *maze) printMaze() {
	fmt.Println(m.antQty)
	for name, room := range m.rooms {
		if room == m.rooms[m.start] {
			fmt.Println("##start")
		}
		if room == m.rooms[m.end] {
			fmt.Println("##end")
		}
		fmt.Printf("%v %d %d\n", name, room.x, room.y)
	}
	visted := []string{}
	for name, room := range m.rooms {
		for _, link := range room.linkTo {
			if isVisited(visted, link) {
				fmt.Printf("%s-%s\n", link, name)
			}
		}
		visted = append(visted, name)
	}
	fmt.Println()
}

func (m *maze) setMaze(fileInput []string) error {
	antQty, err := strconv.Atoi(fileInput[0])
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format, %s is too many / too few ants",
			fileInput[0])
	}
	m.antQty = antQty
	if err := m.setRooms(fileInput); err != nil {
		return err
	}
	if err := m.setLinks(fileInput); err != nil {
		return err
	}
	return nil
}

func (m *maze) setRooms(fileInput []string) error {
	typeOf, entryCnt, i := "", 0, 0

	for ; i < len(fileInput); i++ {
		typeOf = ""
		if fileInput[i] == "##start" || fileInput[i] == "##end" {
			typeOf = fileInput[i][2:]
			i++
		}
		if !strings.Contains(fileInput[i], " ") {
			continue
		}
		name, x, y, isOk := m.setRoom(fileInput[i])
		if !isOk {
			return fmt.Errorf("ERROR: invalid data format. Check room Values: %s", fileInput[i])
		}
		m.rooms[name] = &room{x: x, y: y}

		switch typeOf {
		case "start":
			m.start, entryCnt = name, entryCnt+1
		case "end":
			m.end, entryCnt = name, entryCnt-2
		}
	}
	if entryCnt != -1 {
		return fmt.Errorf("ERROR: invalid data format. Check start/end rooms")
	}
	return nil
}

func (m *maze) setRoom(rmValues string) (string, int, int, bool) {
	values := strings.Fields(rmValues)
	if len(values) != 3 || strings.Contains(rmValues, "-") {
		return "", 0, 0, false
	}
	name := values[0]
	if name[0] == 'L' || name[0] == '#' {
		return "", 0, 0, false
	}
	x, err := strconv.Atoi(values[1])
	if err != nil {
		return "", 0, 0, false
	}
	y, err := strconv.Atoi(values[2])
	if err != nil {
		return "", 0, 0, false
	}
	for key, room := range m.rooms {
		if name == key || (room.x == x && room.y == y) {
			return "", 0, 0, false
		}
	}
	return name, x, y, true
}

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
	}
	return nil
}

func (m *maze) isValidLink(link0, link1 string) bool {
	_, isExist0 := m.rooms[link0]
	room1, isExist1 := m.rooms[link1]
	if !isExist0 || !isExist1 || link0 == link1 {
		return false
	}
	for _, existed := range room1.linkTo {
		if link0 == existed {
			return false
		}
	}
	return true
}
