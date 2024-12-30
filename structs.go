package main

type maze struct {
	antQty   int
	start    string
	end      string
	rooms    map[string]*room
	paths    []path
	solution *solution
}

type room struct {
	x, y   int
	linkTo []string
}

type path struct {
	seq    []string
	length int
}

type solution struct {
	paths  []int
	length int
}
