package main

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestReadInput(t *testing.T) {
	l := Lemz{filename: "test.txt"}
	expected := []string{"3", "##start", "A", "##end", "B"}
	os.WriteFile("./examples/test.txt", []byte("3\n##start\nA\n##end\nB"), 0644)
	defer os.Remove("./examples/test.txt")

	lines, err := l.ReadInput(l.filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], line)
		}
	}
}

func TestAntQty(t *testing.T) {
	l := Lemz{txtLines: []string{"3", "##start", "A", "##end", "B"}}
	expected := 3

	antQty, _ := strconv.Atoi(l.txtLines[0])
	if antQty != expected {
		t.Errorf("Expected %d, got %d", expected, antQty)
	}
}

func TestFileExtension(t *testing.T) {
	l := Lemz{filename: "test.txt"}

	if filepath.Ext(l.filename) != ".txt" {
		t.Errorf("Expected .txt extension, got %s", filepath.Ext(l.filename))
	}
}

func TestFileExistence(t *testing.T) {
	l := Lemz{filename: "test.txt"}
	os.WriteFile("./examples/test.txt", []byte("3\n##start\nA\n##end\nB"), 0644)
	defer os.Remove("./examples/test.txt")

	_, err := os.Stat("./examples/" + l.filename)
	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("Expected file to exist, but it does not")
	}
}

func TestTxtInputOK(t *testing.T) {
	l := Lemz{txtLines: []string{"3", "##start", "A", "##end", "B"}, antQty: 3}

	if !l.TxtInputOK(l.txtLines) {
		t.Errorf("Expected TxtInputOK to return true, got false")
	}
}

func TestHasStartEnd(t *testing.T) {
	l := Lemz{txtLines: []string{"3", "##start", "A", "##end", "B"}}

	if !l.HasStartEnd(l.txtLines) {
		t.Errorf("Expected HasStartEnd to return true, got false")
	}
}
