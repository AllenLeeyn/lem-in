package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: go run . <filename.txt>")
		return
	}
	fileName := os.Args[1]
	fileInput, err := getInput(fileName)
	checkErr(err)

	maze := maze{rooms: make(map[string]*room)}
	err = maze.setMaze(fileInput)
	checkErr(err)

	maze.getPaths([]string{maze.start})
	if len(maze.paths) == 0 {
		checkErr(fmt.Errorf("ERROR: no paths found"))
	}
	maze.getSolution()
	maze.getAntsAssignment()

	maze.printMaze()
	fmt.Println()
	maze.getMoving()
}

func getInput(filename string) ([]string, error) {
	lastIndex := len(filename)
	if lastIndex < 4 || filename[lastIndex-4:] != ".txt" {
		return []string{}, fmt.Errorf("ERROR: only .txt files are allowed")
	}
	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileInput []string
	for scanner.Scan() {
		fileInput = append(fileInput, scanner.Text())
	}
	return fileInput, scanner.Err()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
