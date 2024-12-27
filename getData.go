package main

import (
	"log"
	"strings"
)

func (l *Lemz) GetStartEnd(txtLines []string) (string, string) {
	for i := 0; i < len(txtLines); i++ {
		if txtLines[i] == "##start" && i+1 < len(txtLines) {
			if len(strings.Fields(txtLines[i+1])) != 3 {
				log.Fatal("GetStartEndError: Start room not found.")
			}
			l.start = (strings.Fields(txtLines[i+1]))[0]
		}
		if txtLines[i] == "##end" && i+1 < len(txtLines) {
			if len(strings.Fields(txtLines[i+1])) != 3 {
				log.Fatal("GetStartEndError: End room not found.")
			}
			l.end = (strings.Fields(txtLines[i+1]))[0]
		}
	}
	switch {
	case l.start == "" || l.end == "":
		log.Fatal("GetStartEndError: Start and/or end not found.\n")
	case l.start == l.end:
		log.Fatal("GetStartEndError: Start and end cannot be the same.\n")
	}
	return l.start, l.end
}

// identifies rooms as strings with " " in them
func (l *Lemz) GetRms(txtLines []string) [][3]string {
	for i := 0; i < len(txtLines); i++ {
		for j := 0; j < len(txtLines[i]); j++ {
			if strings.Contains(txtLines[i], " ") {
				rmValues = strings.Fields(txtLines[i])
				l.rmNames = append(l.rmNames, rmValues[0])
				if len(rmValues) == 3 {
					l.rooms = append(l.rooms, [3]string{rmValues[0], rmValues[1], rmValues[2]})
				} else {
					log.Fatal("GetRmsError: Not enough room values.")
				}
				break
			}
		}
	}
	return l.rooms
}

// identifies links as strings with "-" in them
func (l *Lemz) GetLinks(txtLines []string) [][2]string {
	l.graph = make(map[string][]string) // Initialize the graph

	for i := 0; i < len(txtLines); i++ {
		for j := 0; j < len(txtLines[i]); j++ {
			if strings.Contains(txtLines[i], "-") {
				LinkValues := strings.Split(txtLines[i], "-")
				l.links = append(l.links, [2]string{LinkValues[0], LinkValues[1]})

				// Build adjacency list
				l.graph[LinkValues[0]] = append(l.graph[LinkValues[0]], LinkValues[1])
				l.graph[LinkValues[1]] = append(l.graph[LinkValues[1]], LinkValues[0]) // Assuming undirected graph
				break
			}
		}
	}

	// Check if start and end are in the graph
	startFound, endFound := false, false
	for _, link := range l.links {
		if link[0] == l.start || link[1] == l.start {
			startFound = true
		}
		if link[0] == l.end || link[1] == l.end {
			endFound = true
		}
		if link[0] == link[1] {
			log.Fatal("GetLinksError: Link cannot connect to itself.")
		}
	}
	if !startFound {
		log.Fatal("GetLinksError: Start not found in links.")
	}
	if !endFound {
		log.Fatal("GetLinksError: End not found in links.")
	}
	return l.links
}
