package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type Lemz struct {
	filename string
	txtLines []string
	antQty   int
	start    string
	end      string
	rmNames  []string
	rooms    [][3]string
	//linkValues [2]string
	links            [][2]string
	graph            map[string][]string // for adjacency list
	paths            [][]string          // found paths from start to end
	pathsSingleOrder [][]string
	result           string
}

var (
	err        error
	LinkValues [2]string
	rmValues   []string
	aPath      []string
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal(fmt.Errorf("USAGE: Please provide the filename for the TXT file to be read."))
		return
	}

	//defining stuff
	l := Lemz{
		filename: os.Args[1],
	}

	// running funcs to get data from TXT file
	if !l.CheckFile(l.filename) {
		return
	}
	l.txtLines, err = l.ReadInput(l.filename)
	if err != nil {
		log.Fatal(fmt.Errorf("ReadInputError: %s\n", err))
		return
	}
	l.antQty, err = strconv.Atoi(l.txtLines[0])
	if err != nil {
		log.Fatal(fmt.Errorf("antQtyError: %s\n", err))
		return
	}
	if !l.CheckTxt(l.txtLines) {
		return
	}
	l.start, l.end = l.GetStartEnd(l.txtLines)

	l.rooms = l.GetRms(l.txtLines)
	l.links = l.GetLinks(l.txtLines)

	// Run Edmonds-Karp to find paths
	l.EdmondsKarp()

	// Print found paths
	//fmt.Printf("Found paths: %#v\n", l.paths)

	/*fmt.Printf("GetPathSingleOrder(l.links): %#v\n", l.GetPathSingleOrder(l.links))
		l.paths = l.EdmondsKarps(len(l.rooms), l.start, l.end)
	  	fmt.Printf("paths: %#v\n", l.paths) */
	//PRINT FINAL RESULT
	fmt.Printf("Lemz: %#v\n", l)

	// Simulate and print the ant movements
	fmt.Println(strings.Join(l.txtLines, "\n") + "\n")
	l.SimulateAntMovement()

}

// check if filename provided in os.Args is OK
func (l *Lemz) CheckFile(filename string) bool {
	// check TXT file format/ext
	if ext := filepath.Ext(l.filename); ext == "" || ext != ".txt" {
		log.Fatal(fmt.Errorf("CheckFileError: Unsupported file format with extension '%s'. Only TXT files are allowed.\n", filepath.Ext(l.filename)))
		return false
	}
	// check if file even exists
	_, err = os.Stat("./examples/" + l.filename)
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal(fmt.Errorf("CheckFileError: '%s' does not exist.\n", l.filename))
		return false
	}
	return true
}

// read TXT file
func (l *Lemz) ReadInput(filename string) ([]string, error) {
	fileContent, err := os.Open("./examples/" + l.filename)
	if err != nil {
		return nil, fmt.Errorf("OpenError: %s: ", err)
	}
	defer fileContent.Close()

	scanner := bufio.NewScanner(fileContent)
	var fileSlice []string
	for scanner.Scan() {
		fileSlice = append(fileSlice, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("ScannerError: %s.", err)
	}
	return fileSlice, nil
}

func (l *Lemz) CheckTxt(txtLines []string) bool {

	//check quantity of ants
	if l.antQty < 1 || l.antQty > 9223372036854775807 { //too many is > max int64
		log.Fatal(fmt.Errorf("CheckTxt_antQtyError: %d is too many / too few ants.\n", l.antQty))
		return false
	}

	if !slices.Contains(l.txtLines, "##start") || !slices.Contains(l.txtLines, "##end") {
		log.Fatal(fmt.Errorf("CheckTxtError: Provided text has no start and/or end specified.\n"))
		return false
	}
	return true
}
