package main

import (
	"testing"
)

type test struct {
	name         string
	filename     string
	antQty       int
	isThereError bool
}

func TestGetInput(t *testing.T) {
	tests := []test{
		{"ValidFile", "examples/example00.txt", 4, false},
		{"ValidButNoFileName", "tests/.txt", 0, false},
		{"InvalidExtension", "examples/example00.jpg", 0, true},
		{"NonExistentFile", "examples/bad.txt", 0, true},
		{"EmptyFilename", "", 0, true},
		{"DirectoryInsteadOfFile", "examples/", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getInput(tt.filename)
			if (err != nil) != tt.isThereError {
				t.Errorf("func error: '%v' | Expecting test isThereError to be %v.", err, tt.isThereError)
			}
		})
	}
}

func TestSetMaze(t *testing.T) {
	tests := []test{
		{"ValidAntQtyEx01", "examples/example01.txt", 10, false},
		{"ValidAntQtyEx02", "examples/example02.txt", 20, false},
		{"ValidAntQtyEx03", "examples/example03.txt", 4, false},
		{"ValidAntQtyEx04", "examples/example04.txt", 9, false},
		{"ValidAntQtyEx05", "examples/example05.txt", 9, false},
		{"ValidAntQtyEx06", "examples/example06.txt", 100, false},
		{"ValidAntQtyEx07", "examples/example07.txt", 1000, false},
		{"InvalidAntQtyZero", "examples/badexample00.txt", 0, true},
		{"InvalidAntQtyOverMaxInt64", "tests/badexample_overMaxInt64.txt", 0, true},
		{"InvalidAntQtyNotInt", "tests/badexample0a.txt", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInput, err := getInput(tt.filename)
			if err == nil {
				m := maze{antQty: tt.antQty, rooms: make(map[string]*room)}
				err = m.setMaze(fileInput)
				if (err != nil) != tt.isThereError {
					t.Errorf("func error: '%v' | Expecting test isThereError to be %v.", err, tt.isThereError)
				} else if err == nil {
					if m.antQty != tt.antQty {
						t.Errorf("func antQty: '%v' | Expected antQty: %v.", m.antQty, tt.antQty)
					}
				}
			}
		})
	}
}
