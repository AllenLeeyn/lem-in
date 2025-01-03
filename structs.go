package main

type maze struct {
	antQty int
	start  string
	end    string
	rooms  map[string]*room
	paths  []pathStruct
	sol    *solution
	result []string
}

type room struct {
	x, y   int
	linkTo []string
	antID  int
}

type pathStruct struct {
	seq           []string
	length        int
	antsAssigned  int
	antProcessing int
	antsProcessed int
}

type solution struct {
	pathIDs []int
	length  int
}
