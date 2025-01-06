package main

import (
	"testing"
)

type testMaze struct {
	name          string
	filename      string
	start         string
	end           string
	rooms         map[string]*testRoom
	expectedPaths []testPathStruct
}

type testRoom struct {
	linkTo []string
}

type testPathStruct struct {
	seq           []string
	length        int
	antsAssigned  int
	antProcessing int
	antsProcessed int
}

func TestGetPaths(t *testing.T) {
	// Initialize the maze
	tests := []testMaze{
		{
			name:     "example00",
			filename: "examples/example00.txt",
			start:    "0",
			end:      "1",
			rooms: map[string]*testRoom{
				"0": {linkTo: []string{"2"}},
				"2": {linkTo: []string{"0", "3"}},
				"3": {linkTo: []string{"2", "1"}},
				"1": {linkTo: []string{"3"}},
			},
			expectedPaths: []testPathStruct{
				{seq: []string{"2", "3"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
			}},
		{
			name:     "example01",
			filename: "examples/example01.txt",
			start:    "start",
			end:      "end",
			rooms: map[string]*testRoom{
				"start": {linkTo: []string{"t", "h", "0"}},
				"t":     {linkTo: []string{"E", "start"}},
				"h":     {linkTo: []string{"A", "n", "start"}},
				"0":     {linkTo: []string{"o", "start"}},
				"o":     {linkTo: []string{"n", "0"}},
				"n":     {linkTo: []string{"e", "m", "h", "o"}},
				"e":     {linkTo: []string{"end", "n"}},
				"a":     {linkTo: []string{"m", "E"}},
				"A":     {linkTo: []string{"h", "c"}},
				"c":     {linkTo: []string{"k", "A"}},
				"k":     {linkTo: []string{"end", "c"}},
				"E":     {linkTo: []string{"a", "t"}},
				"m":     {linkTo: []string{"end", "n", "a"}},
				"end":   {linkTo: []string{"e", "k", "m"}},
			},
			expectedPaths: []testPathStruct{
				{seq: []string{"t", "E", "a", "m"}, length: 5, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"t", "E", "a", "m", "n", "e"}, length: 7, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"t", "E", "a", "m", "n", "h", "A", "c", "k"}, length: 10, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"h", "A", "c", "k"}, length: 5, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"h", "n", "e"}, length: 4, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"h", "n", "m"}, length: 4, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"0", "o", "n", "e"}, length: 5, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"0", "o", "n", "m"}, length: 5, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"0", "o", "n", "h", "A", "c", "k"}, length: 8, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
			}},
		/* {
			name:     "example02",
			filename: "examples/example02.txt",
			start:    "0",
			end:      "3",
			rooms: map[string]*testRoom{
				"0": {linkTo: []string{"1", "3"}},
				"1": {linkTo: []string{"0", "2"}},
				"2": {linkTo: []string{"1", "3"}},
				"3": {linkTo: []string{"0", "2"}},
			},
			expectedPaths: []testPathStruct{
				{seq: []string{"1", "2"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
			},
		}, */
		{
			name:     "example03",
			filename: "examples/example03.txt",
			start:    "0",
			end:      "5",
			rooms: map[string]*testRoom{
				"0": {linkTo: []string{"1", "2", "3"}},
				"1": {linkTo: []string{"0", "4"}},
				"2": {linkTo: []string{"4", "0"}},
				"3": {linkTo: []string{"0", "4"}},
				"4": {linkTo: []string{"2", "1", "5", "3"}},
				"5": {linkTo: []string{"4"}},
			},
			expectedPaths: []testPathStruct{
				{seq: []string{"1", "4"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"2", "4"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"3", "4"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
			},
		},
		{
			name:     "example04",
			filename: "examples/example04.txt",
			start:    "richard",
			end:      "peter",
			rooms: map[string]*testRoom{
				"richard":  {linkTo: []string{"dinish", "gilfoyle", "erlich"}},
				"dinish":   {linkTo: []string{"richard", "jimYoung"}},
				"jimYoung": {linkTo: []string{"dinish", "erlich", "peter"}},
				"gilfoyle": {linkTo: []string{"richard", "peter", "erlich"}},
				"peter":    {linkTo: []string{"gilfoyle", "jimYoung"}},
				"erlich":   {linkTo: []string{"gilfoyle", "richard", "jimYoung"}},
			},
			expectedPaths: []testPathStruct{
				{seq: []string{"dinish", "jimYoung", "erlich", "gilfoyle"}, length: 5, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"dinish", "jimYoung"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"gilfoyle"}, length: 2, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"gilfoyle", "erlich", "jimYoung"}, length: 4, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"erlich", "gilfoyle"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"erlich", "jimYoung"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInput, err := getInput(tt.filename)
			if err == nil {
				m := maze{rooms: make(map[string]*room), start: tt.start, end: tt.end}
				err = m.setMaze(fileInput)
				m.getPaths([]string{m.start})

				if len(m.paths) != len(tt.expectedPaths) {
					t.Errorf("Expected paths length %d, but got %d", len(tt.expectedPaths), len(m.paths))
				} else {
					for i := range m.paths {
						if len(m.paths[i].seq) != len(tt.expectedPaths[i].seq) || m.paths[i].length != tt.expectedPaths[i].length {
							t.Errorf("Expected path at index %d to be %v, but got %v", i, tt.expectedPaths[i], m.paths[i])
							continue
						}
						for j := range m.paths[i].seq {
							if m.paths[i].seq[j] != tt.expectedPaths[i].seq[j] {
								t.Errorf("Expected path at index %d, element %d to be %v, but got %v", i, j, tt.expectedPaths[i].seq[j], m.paths[i].seq[j])
							}
						}
					}
				}
			}
		})
	}
}