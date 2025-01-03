package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: go run . <filename.txt>")
		return
	}
	fileInput, err := getInput(os.Args[1])
	checkErr(err)

	m := maze{rooms: make(map[string]*room)}
	err = m.setMaze(fileInput)
	checkErr(err)

	m.getPaths([]string{m.start})
	if len(m.paths) == 0 {
		checkErr(fmt.Errorf("ERROR: no paths found"))
	}

	m.getSolution()
	m.getAntsAssignment()
	m.printMaze()
	m.getMoving()
}

// getInput() reads contents of .txt file. Ignore empty newlines and comments
func getInput(filename string) ([]string, error) {
	if !strings.HasSuffix(filename, ".txt") {
		return nil, fmt.Errorf("ERROR: only .txt files are allowed")
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileInput []string
	for scanner.Scan() {
		fileLine := scanner.Text()
		if strings.HasPrefix(fileLine, "#") && !strings.HasPrefix(fileLine, "##") {
			continue
		}
		fileInput = append(fileInput, fileLine)
	}
	return fileInput, scanner.Err()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
