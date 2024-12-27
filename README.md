# lem-in

## Objective

This `lem-in` project is a digital version of an ant farm. The goals of the program are as follows:

- read file describing the ants and the colony
- find the quickest path to get _n_ ants across a colony (avoid traffic jams)
- only one ant can move through a link at one time (no walking over other ants)
- display the following:
  - content of the file passed as argument (make an ant farm with tunnels and rooms)
  - each move the ant makes from room to room (look at how they find the exit)
  - standard output display format:

    ```bash
    number_of_ants
    the_rooms
    the_links

    Lx-y Lz-w Lr-o ... // ants move at each turn
    ```

    - x, z, r represent the ant number
    - y, w, o represent the room names
    - rooms are defined by `name coord_x coord_y`
    - coordinates of the rooms are always in `int`
    - links are defined by `name1-name2`

- all ants start in the room ##start
- bring them to the room ##end as few moves as possible

- ERROR: invalid data format
  - rooms that link to themselves
  - too many/ too few ants
  - no ##start/ ##end
  - duplicated rooms
  - link to unknown rooms
  - room with invalid coordinates
  - a variety of invalid or poorly formatted input
  - A room will never start with the letter _L_ or with _#_ and must have no spaces
  - two rooms can't have more than one tunnel connecting them
  - each room can contain only one ant at a time

## Task

- read file (Jedi)
- validate input (Jedi)
- get number of ants (Jedi)
- get rooms (Jedi)
- map out rooms and links/tunnels (?)

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
    Start{start} ==> h([h]) ==> A([A]) ==> c([c]) ==> k([k]) ==> End[end]
    h([h]) -.-> n([n])
    n([n]) -.-> m([m])
    Start{start} ==> 0([0]) ==> o([o]) ==> n([n]) ==> e([e]) ==> End[end]

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
Start{start} ==> A0([A0]) ==> A1([A1]) ==> A2([A2]) ==> End[end]
Start{start} ==> C0([C0]) ==> C1([C1]) ==> C2([C2]) ==> C3([C3]) ==> I4([I4]) ==> I5([I5]) ==> End[end]
Start{start} -.-> G1([G1]) -.-> G2([G2]) -.-> G3([G3]) -.-> G4([G4]) -.-> D3([D3])
Start{start} ==> B0([B0]) ==> B1([B1]) ==> E2([E2]) ==> D2([D2]) ==> D3([D3]) ==> End[end]
A0([A0]) -.-> D1([D1]) -.-> F2([F2]) -.-> F3([F3]) -.-> F4([F4]) -.-> End[end]
A1([A1]) -.-> B1([B1])
A2([A2]) -.-> C3([C3])
D2([D2]) -.-> F3([F3])
F2([F2]) -.-> H3([H3]) -.-> H4([H4])

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