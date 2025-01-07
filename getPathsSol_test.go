package main

import (
	"strings"
	"testing"
)

type testMaze struct {
	name          string
	filename      string
	start         string
	end           string
	rooms         map[string]*testRoom
	expectedPaths []testPathStruct
	sol           *testSolution
	expectedSeq   [][]string
	expectedError bool
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

type testSolution struct {
	pathIDs []int
	length  int
}

func TestGetPaths(t *testing.T) {
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
			expectedPaths: []testPathStruct{
				{seq: []string{"1", "2"}, length: 3, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
				{seq: []string{}, length: 1, antsAssigned: 0, antProcessing: 0, antsProcessed: 0},
			},
		},
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

func TestGetPaths_Duplicates(t *testing.T) {
	tests := []testMaze{
		{
			name:          "example00",
			filename:      "examples/example00.txt",
			expectedError: false,
		},
		{
			name:          "example01",
			filename:      "examples/example01.txt",
			expectedError: false,
		},
		{
			name:          "example02",
			filename:      "examples/example02.txt",
			expectedError: false,
		},
		{
			name:          "example03",
			filename:      "examples/example03.txt",
			expectedError: false,
		},
		{
			name:          "example04",
			filename:      "examples/example04.txt",
			expectedError: false,
		},
		{
			name:          "example05",
			filename:      "examples/example05.txt",
			expectedError: false,
		},
		{
			name:     "with duplicates",
			filename: "",
			start:    "start",
			end:      "end",
			rooms: map[string]*testRoom{
				"start": {linkTo: []string{"A", "B"}},
				"A":     {linkTo: []string{"C"}},
				"B":     {linkTo: []string{"C"}},
				"C":     {linkTo: []string{"end"}},
				"end":   {linkTo: []string{}}},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInput, err := getInput(tt.filename)
			if err == nil {
				tt.expectedPaths = []testPathStruct{} // Reset paths
				m := maze{rooms: make(map[string]*room), start: tt.start, end: tt.end}
				err = m.setMaze(fileInput)
				m.getPaths([]string{m.start})

				seen := make(map[string]int)
				hasDuplicates := false
				if tt.filename != "" {
					for _, path := range m.paths {
						serialized := serializeSlice(path.seq)
						if _, exists := seen[serialized]; exists {
							hasDuplicates = true
							break
						}
						seen[serialized]++
					}

					if hasDuplicates != tt.expectedError {
						t.Errorf("Test %q failed: expected duplicates=%v, but got %v", tt.name, tt.expectedError, hasDuplicates)
					}
				} else {
					for _, path := range tt.expectedPaths {
						serialized := serializeSlice(path.seq)
						if _, exists := seen[serialized]; exists {
							hasDuplicates = true
							break
						}
						seen[serialized]++
					}

					if hasDuplicates != tt.expectedError {
						t.Errorf("Test %q failed: expected duplicates=%v, but got %v", tt.name, tt.expectedError, hasDuplicates)
					}
				}
			}
		})
	}
}

// serializeSlice converts a slice of strings into a single string for map key usage.
func serializeSlice(slice []string) string {
	return strings.Join(slice, ",")
}

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
