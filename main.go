package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type lemz struct {
	filename string
	txtLines []string
	line     string
	antQty   int
	start    string
	end      string
	rmNames  []string
	paths    []string
	err      error
	result   []string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal(fmt.Errorf("USAGE: Please provide the filename for the TXT file to be read."))
		return
	}

	//defining stuff
	l := lemz{
		filename: os.Args[1],
		result:   []string{"||tempResult||"}, //replace with final result later
	}
	l.txtLines = func() []string {
		txtLns, err := l.ReadInput(l.filename)
		if err != nil {
			log.Fatal(fmt.Errorf("ReadInputError: %s\n", err))
		}
		return txtLns
	}()
	l.antQty = func() int {
		antQt, err := strconv.Atoi(l.txtLines[0])
		if err != nil {
			log.Fatal(fmt.Errorf("antQtyError: %s\n", err))
		}
		return antQt
	}()

	// check TXT file format/ext
	if filepath.Ext(l.filename) == "" {
		log.Fatal(fmt.Errorf("Error: Filename provided has no file extension. Only TXT files are allowed."))
		return
	}
	if filepath.Ext(l.filename) != ".txt" {
		log.Fatal(fmt.Errorf("Error: Unsupported file format with extension '%s'. Only TXT files are allowed.\n", filepath.Ext(l.filename)))
		return
	}

	// check if file even exists
	_, err := os.Stat("./examples/" + l.filename)
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal(fmt.Errorf("Error: '%s' does not exist.\n", l.filename))
		return
	}

	//PRINT FINAL RESULT
	if l.TxtInputOK(l.txtLines) {
		fmt.Println(strings.Join(l.txtLines, "\n") + "\n" + l.result[0])
	}
}

// read TXT file
func (l *lemz) ReadInput(filename string) ([]string, error) {
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
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ScannerError: %s.", err)
	}
	return fileSlice, nil
}

func (l *lemz) TxtInputOK(txtLines []string) bool {

	//check quantity of ants
	if l.antQty < 1 || l.antQty > 9223372036854775807 { //too many is > max int
		log.Fatal(fmt.Errorf("antQtyError: %d is too many / too few ants.\n", l.antQty))
		return false
	}

	if !l.HasStartEnd(l.txtLines) {
		log.Fatal(fmt.Errorf("InputError: Input has no start and/or end specified.\n"))
		return false
	}

	return true
}

// To be cont: below doesnt work
func (l *lemz) HasStartEnd([]string) bool {
	for _, ln := range l.txtLines {
		if ln == "##start" || ln == "##end" {
			return true
		}
	}
	return false
}
