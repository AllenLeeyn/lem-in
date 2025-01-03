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

const resetColor = "\033[00m"

var colors = []string{
	"\033[0;31m", //red
	"\033[0;32m", //green
	"\033[0;33m", //yellow
	"\033[0;34m", //blue
	"\033[0;35m", //magenta
	"\033[0;36m", //cyan
	"\033[1;91m", //high intense red
	"\033[1;92m", //high intense green
	"\033[1;93m", //high intense yellow
	"\033[1;94m", //high intense blue
	"\033[1;95m", //high intense magenta
	"\033[1;96m", //high intense cyan
}
