package main

type maze struct {
	antQty       int
	start        string
	end          string
	rooms        map[string]*room
	paths        []pathStruct
	optimalPaths []int
	inputLines   string // sanitize lines of input
	movement     string
}

type room struct {
	x, y   int
	linkTo []string
	antID  int
}

type pathStruct struct {
	path          []string
	length        int
	antsSet       int
	antProcessing int
	antsProcessed int
}

const resetColor = "\033[00m"

var colors = []string{
	"\033[31m",   //red
	"\033[32m",   //green
	"\033[33m",   //yellow
	"\033[34m",   //blue
	"\033[35m",   //magenta
	"\033[36m",   //cyan
	"\033[1;91m", //high intense red
	"\033[1;92m", //high intense green
	"\033[1;93m", //high intense yellow
	"\033[1;94m", //high intense blue
	"\033[1;95m", //high intense magenta
	"\033[1;96m", //high intense cyan
}
