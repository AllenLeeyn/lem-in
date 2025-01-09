# lem-in

`lem-in` is a digital version of an ant farm. This program takes a `TXT` file that describes the ants, rooms, and links. The program then tries to find the most optimal set of paths that takes the ants from the start room to the end room in the shortest possible total turns.

## Installation

1. Make sure the following are installed in your sytem before running the program:

- [Go](https://golang.org/doc/install) (version 1.20 or later recommended)
- Git, such as [Gitea](https://01.gritlab.ax/git/), to clone the repository


2. Clone the repository by running the following code:
	
	```bash
	https://01.gritlab.ax/git/ylee/lem-in.git
	cd lem-in
	```

## Usage

You can now use the application by running the following command:

```bash
go run . <filename.txt>
```

Make sure to replace `<filename.txt>` with the specific filename of the `TXT` file you want to access. Sample `TXT` files are provided inside the `examples` directory. For example, you may access one of the examples as follows:

```bash
go run . examples/example00.txt
```

You should then see the output displayed on the terminal in the following format:

```bash
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```

- `x`, `z`, and `r` represent the ant numbers (going from 1 to `number_of_ants`)
- `y`, `w`, and `o` represent the room names

## Input Validity and Error Handling

- Only `TXT` files are accepted.
- The first line is considered as the number of ants and, therefore, should always be a positive integer.
- The rooms should be in the format `roomName coord_x coord_y`, and `coord_x coord_y` should always be integers; for example, `start 1 2`.
- The room names should not start with the letter `L` or with the symbol `#` and must have no spaces within it.
- There should be only one `##start` and one `##end` commands and corresponding rooms.
- The start and end rooms should not be the same.
- The links should be in the format `roomName1-roomName2`; for example, `1-2`.
- The rooms cannot link to itself; for example, `3-3` is invalid.
- A link can join two rooms only once; for example, having both `1-2` and `2-1` as links is invalid.
- Any unknown command that starts with `#` will be ignored and not included in the final output.
- Except for the start and end rooms, a room can only have one ant at a time per turn, ants cannot occupy the same room in a turn, and a link can be used only once per turn.
- An ant can move only once per turn, and the destination room should be empty.
- Unit tests are provided.

## Implementation

Here is a flowchart visually describing this `lem-in` application:

```mermaid
flowchart LR

main["take filename/os.Args[1]"] ==> getInput ==> setMaze ==> getPaths ==> getSolution ==> getAntsAssignment ==> printMaze ==> getMoving ==> printMovement[print final movement]

%% main.go > func getInput
  subgraph getInput[func getInput]
      chkExt([check file ext]) --> open([open file]) --> scan([scan file contents]) --> sanitize([sanitize input]) --> fileIn([create slice of strings])
  end

%% setMaze.go
subgraph setMaze[func setMaze]
direction LR
setAntQty ==> setRooms ==> setLinks

    subgraph setAntQty[func setAntQty]
    chkAnt([validate if antQty is int])
    end

    subgraph setRooms[func setRooms]
    chkStartEnd([check ##start and ##end]) --> sp([process only lines that contain space]) --> setRoom

        subgraph setRoom[func setRoom]
        direction TB
        splitS([split line using space as separator]) --> chkValues([validate if line has 3 values, doesn't contain a dash '-'])
        end

    end

    subgraph setLinks[func setLinks]
    dash([process only lines that contain a dash: '-']) --> splitL([split line and save link0, link1]) --> isValidLink --> adjL([populate adjacency list m.rooms.linkTo: list of all rooms and all its links, considering both link0 and link1])

        subgraph isValidLink[func isValidLink]
        chkRmExists([check room exists]) --> chkLinkToItself([check that room doesn't link to itself]) --> chkSameLink([link can't be more than 1])
        end

    end

end

%% getPaths.go
subgraph getPaths[func getPaths]
chkAllLinks([check all links to curRoom]) --> isVisited --> recursivGetPath([find next room by calling getPaths recursively]) --> pathStruct([populate pathStruct once end room is reached])

    subgraph isVisited[func !isVisited: check room has not been visited in curPath]
    end

end

%% getSolution.go
subgraph getSolution[func getSolution]
getFlow([get max flow]) --> searchSolution --> appendEnd([append end room to path found]) --> sortSolution

    subgraph searchSolution[func searchSolution]
    direction TB
    isPathsClash --> recursivSearchSol([recursively search for paths until max flow is reached]) --> solStruct([populate solution struct])
    
        subgraph isPathsClash[func isPathsClash: check if room is at the same position in any path]
        end

    end

    subgraph sortSolution[func sortSolution: shortest to longest path length]
    end

end

%% getMoving.go
subgraph getAntsAssignment[func getAntsAssignment]
iterateSolPath([iterate thru all path solutions]) --> assignAnts([distribute ants based on path length])
end

subgraph printMaze[func printMaze]
printInput([print sanitized input + empty line])
end

subgraph getMoving[func getMoving]
iterateSolPathRms([take each ant thru each room of each path solution]) --> getMove

    subgraph getMove["func getMove: print each ant movement (ant-room) and set colors"]
    end

end
```

The following are flowcharts to visualize some of the provided examples:

#### example00

```mermaid
flowchart LR
    Start{0} ==> 2([2]) ==> 3([3]) ==> End[1]

%%style
    style Start fill:green
    style End fill:red
```

#### example01

```mermaid
flowchart LR
    Start{start} ==> t([t]) ==> E([E]) ==> a([a]) ==> m([m]) ==> End[end]
    Start{start} ==> 0([0]) ==> o([o]) ==> n([n]) ==> e([e]) ==> End[end]
    Start{start} ==> h([h]) ==> A([A]) ==> c([c]) ==> k([k]) ==> End[end]
    h([h]) -.-> n([n])
    n([n]) -.-> m([m])

%%style
    style Start fill:green
    style End fill:red
```

#### example02

```mermaid
flowchart LR
    Start{0} ==> 1([1]) ==> 2([2]) ==> End[3]
    Start{0} ==> End[3]
    
%%style
    style Start fill:green
    style End fill:red
```

#### example03

```mermaid
flowchart LR
    Start{0} -.-> 2([2]) -.-> 4([4])
    Start{0} ==> 1([1]) ==> 4([4]) ==> End[5]
    Start{0} -.-> 3([3]) -.-> 4([4])

%%style
    style Start fill:green
    style End fill:red
```

#### example04/06/07

```mermaid
flowchart LR
    Start{richard} ==> d([dinish]) ==> j([jimYoung]) ==> End[peter]
    Start{richard} ==> g([gilfoyle]) ==> End[peter]
    Start{richard} -.-> e([erlich]) -.-> j([jimYoung])
    g([gilfoyle]) -.-> e([erlich])

%%style
    style Start fill:green
    style End fill:red
```

#### example05

```mermaid
flowchart LR
Start{start} ==> B0([B0]) ==> B1([B1]) ==> E2([E2]) ==> D2([D2]) ==> F3([F3]) ==> F4([F4]) ==> End[end]
Start{start} ==> A0([A0]) ==> A1([A1]) ====> A2([A2]) ==> End[end]
A2([A2]) -.-> C3([C3])
Start{start} ==> C0([C0]) ==> C1([C1]) ==> C2([C2]) ====> C3([C3]) ==> I4([I4]) ==> I5([I5]) ==> End[end]

A0([A0]) -.-> D1([D1]) -.-> F2([F2]) -.-> H3([H3]) -.-> H4([H4])
Start{start} ==> G0([G0]) ==> G1([G1]) ==> G2([G2]) ==> G3([G3]) ==> G4([G4]) ==> D3([D3]) ==> End[end]
A1([A1]) -.-> B1([B1])

%%style
    style Start fill:green
    style End fill:red
```

#### badexample01

```mermaid
flowchart LR
 subgraph start["start and links only"]
        
    Start{"0"} --> 1(["1"]) & 5(["5"]) & 9(["9"]) & 10(["10"])
    1 --> 2(["2"]) & 11(["11"])
    2 --> 3(["3"])
    3 --> 3
    5 --> 6(["6"])
    6 --> 7(["7"])
    7 --> 8(["8"])
    8 --> 7
    9 --> 2
    10 --> 16(["16"])
    16 --> 7
    11 --> 12(["12"])
    12 --> 13(["13"])
    13 --> 14(["14"])
    14 --> 15(["15"])
    15 --> 1
end

subgraph end["end doesn't link anywhere"]
    End["4"]
end

%%style
    style Start fill:green
    style End fill:red
```

## Members

* [Inka Säävuori](https://github.com/Inkasaa) 👑
* [Jedi Reston](https://github.com/jeeeeedi) 🤓
* [Yuanneng Lee (Allen)](https://github.com/AllenLeeyn) 🤖