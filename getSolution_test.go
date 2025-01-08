package main

import (
	"fmt"
	"testing"
)

func TestGetSolution(t *testing.T) {
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
			expectedSeq: [][]string{
				{"2", "3", "1"},
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
			expectedSeq: [][]string{
				{"t", "E", "a", "m", "end"},
				{"h", "A", "c", "k", "end"},
				{"0", "o", "n", "e", "end"},
			}},
		{
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
			expectedSeq: [][]string{
				{"3"},
				{"1", "2", "3"},
			}},
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
			expectedSeq: [][]string{
				{"1", "4", "5"},
			}},
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
			expectedSeq: [][]string{
				{"gilfoyle", "peter"},
				{"dinish", "jimYoung", "peter"},
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInput, err := getInput(tt.filename)
			if err != nil {
				t.Fatalf("Failed to get input: %v", err)
			}

			m := maze{rooms: make(map[string]*room), start: tt.start, end: tt.end}
			err = m.setMaze(fileInput)
			if err != nil {
				t.Fatalf("Failed to set maze: %v", err)
			}

			m.getPaths([]string{m.start})
			m.getSolution()

			if m.sol == nil {
				t.Fatalf("Expected solution, but got nil")
			}

			if len(m.sol.pathIDs) != len(tt.expectedSeq) {
				t.Errorf("Expected pathIDs length %d, but got %d", len(tt.expectedSeq), len(m.sol.pathIDs))
			}

			for i, id := range m.sol.pathIDs {
				if !compareSlices(m.paths[id].seq, tt.expectedSeq[i]) {
					t.Errorf("Expected path sequence at index %d to be %v, but got %v", i, tt.expectedSeq[i], m.paths[id].seq)
				}
			}
		})
	}
}

// compareSlices compares two slices of strings for equality
func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGetSolution_Duplicates(t *testing.T) {
	tests := []testMaze{
		{
			name:         "example00",
			filename:     "examples/example00.txt",
			isThereError: false,
		},
		{
			name:         "example01",
			filename:     "examples/example01.txt",
			isThereError: false,
		},
		{
			name:         "example02",
			filename:     "examples/example02.txt",
			isThereError: false,
		},
		{
			name:         "example03",
			filename:     "examples/example03.txt",
			isThereError: false,
		},
		{
			name:         "example04",
			filename:     "examples/example04.txt",
			isThereError: false,
		},
		{
			name:         "example05",
			filename:     "examples/example05.txt",
			isThereError: false,
		},
		{
			name:     "exampleWithDupeSeq",
			filename: "",
			start:    "start",
			end:      "end",
			/* rooms: map[string]*testRoom{
			"start": {linkTo: []string{"A", "B"}},
			"A":     {linkTo: []string{"C"}},
			"B":     {linkTo: []string{"C"}},
			"C":     {linkTo: []string{"end"}},
			"end":   {linkTo: []string{}}},
			expectedPaths: []testPathStruct{
				{seq: []string{"A", "C"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"A", "C"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{"B", "C"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
			}, */
			expectedSeq: [][]string{
				{"A", "C"},
				{"A", "C"},
				{"B", "C"},
			},
			isThereError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var m maze

			if tt.filename != "" { // if filename is provided, then use the input from that file
				// Load input from file
				fileInput, err := getInput(tt.filename)
				if err != nil {
					t.Fatalf("Failed to read input file %q: %v", tt.filename, err)
				}
				m = maze{rooms: make(map[string]*room)}
				err = m.setMaze(fileInput)
				if err != nil {
					t.Fatalf("Failed to set maze: %v", err)
				}

				m.getPaths([]string{m.start})
				m.getSolution()

				// Check for duplicates in `m.paths[id].seq`
				seen := make(map[string]int)
				hasDuplicates := false
				duplicateDetails := []string{}

				for i, id := range m.sol.pathIDs {
					serialized := serializeSlice(m.paths[id].seq)
					if index, exists := seen[serialized]; exists {
						hasDuplicates = true
						duplicateDetails = append(duplicateDetails,
							fmt.Sprintf("Repeated seqs found at index %d %v and index %d %v", index, m.paths[index].seq, i, m.paths[id].seq))
					}
					seen[serialized] = i
				}

				if hasDuplicates != tt.isThereError {
					t.Errorf("Test %q failed: expected duplicates=%v, but got %v. Details: %v",
						tt.name, tt.isThereError, hasDuplicates, duplicateDetails)
				}
			} else { // if no filename is provided, then use the provied test expectedSeqs
				// Check for duplicates in `tt.expectedSeqs`
				seen := make(map[string]int)
				hasDuplicates := false
				duplicateDetails := []string{}

				for i, seq := range tt.expectedSeq {
					serialized := serializeSlice(seq)
					if index, exists := seen[serialized]; exists {
						hasDuplicates = true
						duplicateDetails = append(duplicateDetails,
							fmt.Sprintf("Repeated seqs found at index %d %v and index %d %v", index, tt.expectedSeq[index], i, seq))
					}
					seen[serialized] = i
				}

				if hasDuplicates != tt.isThereError {
					t.Errorf("Test %q failed: expected duplicates=%v, but got %v. Details: %v",
						tt.name, tt.isThereError, hasDuplicates, duplicateDetails)
				}
			}
		})
	}
}
