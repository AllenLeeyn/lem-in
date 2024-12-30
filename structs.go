package main

type maze struct {
	antQty int
	start  string
	end    string
	rooms  map[string]*room
	paths  []path
	sol    *solution
}

type room struct {
	x, y   int
	linkTo []string
	antNm  int
}

type path struct {
	seq    []string
	length int
	antCnt int
}

type solution struct {
	paths  []int
	length int
}
