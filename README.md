# LEM-IN

## Objective
Code a digital version of an ant farm

- read file describing the ants and the colony
- find the quickest path to get n ants across a colony (avoid traffic jams)
- only one ant can move through a link at one time (no walking over other ants)

- display:
    - content of the file passed as argument (make an ant farm with tunnels and rooms)
    - each move the ant make from room to room (look at how they find the exit)
    - standard output display format:
        ```
        number_of_ants
        the_rooms
        the_links

        Lx-y Lz-w Lr-o ... // ants move at each turn
        ```
        - x,z, r represents the ant number
        - y,w, o represents the rooms names
        - rooms are defined by `name coord_x coord_y`
         - coordinates of room will be int
        - links are defined by `name1-name2`

- all ants start in the room ##start
- bring them to the room ##end as few moves as possible

- ERROR: invalid data format
    - rooms that link to themselves
    - too many/ too few ants
    - no ##start/ ##end
    - duplicated rooms
    - link to unkwnons rooms
    - room with invlid coorediates
    - a variety of invalid or poorly formatted input
    - A room will never start with the letter L or with # and must have no spaces
    - two rooms can't have more than one tunnel connecting them
    - each room can contain only one ant at a time

## Task
- read file
- validate input
 - get number of ants
 - get rooms
- map out rooms and links/ tunnels (?)

#### example00
```mermaid
flowchart LR
start0---2---3---end1
```
#### example01
```mermaid
flowchart LR
start---t & 0 & h
t---E---a---m---'end'
0---'o'---n---e---'end'
h---n---m
h---A---c---k---'end'
```

#### example02
```mermaid
flowchart LR
start0---1
start0---end3
1---2
2<---end3
```

#### example03
```mermaid
flowchart LR
start0-->1
2<-->4
1-->4
start0-->2
4-->end5
start0-->3
3-->4
```

#### example04
```mermaid
flowchart LR
richardStart-->erlich-->peterExit & jimYoung
richardStart-->gilfoyle-->peterExit & erlich
richardStart-->dinish-->jimYoung-->peterExit
```

#### example05
```mermaid
flowchart LR
start<-->A0 & B0 & C0 & G0
B0<-->B1
A0<-->A1
A0<-->D1
A1<-->A2
A1<-->B1
A2<-->'end'
A2<-->C3
B1<-->E2
C0<-->C1
C1<-->C2
C2<-->C3
C3<-->I4
D1<-->D2
D1<-->F2
D2<-->E2
D2<-->D3
D2<-->F3
D3<-->'end'
F2<-->F3
F3<-->F4
F4<-->'end'
G0<-->G1
G1<-->G2
G2<-->G3
G3<-->G4
G4<-->D3
H3<-->F2
H3<-->H4
H4<-->A2
I4<-->I5
I5<-->'end'
```
