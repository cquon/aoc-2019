package main

import (
	"os"
	"log"
	"math"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

/*
--- Day 3: Crossed Wires ---
The gravity assist was successful, and you're well on your way to the Venus refuelling station. During the rush back on Earth, the fuel management system wasn't completely installed, so that's next on the priority list.

Opening the front panel reveals a jumble of wires. Specifically, two wires are connected to a central port and extend outward on a grid. You trace the path each wire takes as it leaves the central port, one wire per line of text (your puzzle input).

The wires twist and turn, but the two wires occasionally cross paths. To fix the circuit, you need to find the intersection point closest to the central port. Because the wires are on a grid, use the Manhattan distance for this measurement. While the wires do technically cross right at the central port where they both start, this point does not count, nor does a wire count as crossing with itself.

For example, if the first wire's path is R8,U5,L5,D3, then starting from the central port (o), it goes right 8, up 5, left 5, and finally down 3:

...........
...........
...........
....+----+.
....|....|.
....|....|.
....|....|.
.........|.
.o-------+.
...........
Then, if the second wire's path is U7,R6,D4,L4, it goes up 7, right 6, down 4, and left 4:

...........
.+-----+...
.|.....|...
.|..+--X-+.
.|..|..|.|.
.|.-X--+.|.
.|..|....|.
.|.......|.
.o-------+.
...........
These wires cross at two locations (marked X), but the lower-left one is closer to the central port: its distance is 3 + 3 = 6.

Here are a few more examples:

R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83 = distance 159
R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135
What is the Manhattan distance from the central port to the closest intersection?
*/

func readWireInputs(fileName string) ([]string, []string) {
	var wire1 []string
	var wire2 []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read first line
	scanner.Scan()
	wireLine1 := scanner.Text()
	wire1 = strings.Split(wireLine1, ",")

	// Read second line
	scanner.Scan()
	wireLine2 := scanner.Text()
	wire2 = strings.Split(wireLine2, ",")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wire1, wire2
}

func populateCoordinates(wireCoordinates map[int]map[int]struct{}, wireInput []string) {
	x := 0
	y := 0
	for _, directionInput := range wireInput {
		direction := string(directionInput[0])
		number, err := strconv.Atoi(strings.TrimPrefix(directionInput, direction))
		if err != nil {
			panic("Input string must have L|D|U|R followed by a numeric value")
		}
		switch direction {
		case "L":
			for i:=0;i<number;i++ {
				if wireCoordinates[x-1] == nil {
					wireCoordinates[x-1] = make(map[int]struct{})
				}
				wireCoordinates[x-1][y] = struct{}{}
				x -= 1
			}
		case "D":
			for i:=0;i<number;i++ {
				if wireCoordinates[x] == nil {
					wireCoordinates[x] = make(map[int]struct{})
				}
				wireCoordinates[x][y-1] = struct{}{}
				y -= 1
			}
		case "U":
			for i:=0;i<number;i++ {
				if wireCoordinates[x] == nil {
					wireCoordinates[x] = make(map[int]struct{})
				}
				wireCoordinates[x][y+1] = struct{}{}
				y += 1
			}
		case "R":
			for i:=0;i<number;i++ {
				if wireCoordinates[x+1] == nil {
					wireCoordinates[x+1] = make(map[int]struct{})
				}
				wireCoordinates[x+1][y] = struct{}{}
				x += 1
			}
		}
	}
}

func distance(x1, y1, x2, y2 int) int {
	xDistance := x2 - x1
	yDistance := y2 - y1

	if xDistance < 0 {
		xDistance *= -1
	}

	if yDistance < 0 {
		yDistance *= -1
	}

	return xDistance + yDistance
}

func getSmallestDistance(wire1Coordinates map[int]map[int]struct{}, wireInput []string) int {
	x := 0
	y := 0
	minDistance := math.MaxInt32
	for _, directionInput := range wireInput {
		direction := string(directionInput[0])
		number, err := strconv.Atoi(strings.TrimPrefix(directionInput, direction))
		if err != nil {
			panic("Input string must have L|D|U|R followed by a numeric value")
		}
		switch direction {
		case "L":
			for i:=0;i<number;i++ {
				if wire1Coordinates[x-1] != nil {
					if _, exists := wire1Coordinates[x-1][y]; exists {
						dist := distance(0, 0, x-1, y)
						if dist < minDistance {
							minDistance = dist
						}
					}
				}
				x -= 1
			}
		case "D":
			for i:=0;i<number;i++ {
				if wire1Coordinates[x] != nil {
					if _, exists := wire1Coordinates[x][y-1]; exists {
						dist := distance(0, 0, x, y-1)
						if dist < minDistance {
							minDistance = dist
						}
					}
				}
				y -= 1
			}
		case "U":
			for i:=0;i<number;i++ {
				if wire1Coordinates[x] != nil {
					if _, exists := wire1Coordinates[x][y+1]; exists {
						dist := distance(0, 0, x, y+1)
						if dist < minDistance {
							minDistance = dist
						}
					}
				}
				y += 1
			}
		case "R":
			for i:=0;i<number;i++ {
				if wire1Coordinates[x+1] != nil {
					if _, exists := wire1Coordinates[x+1][y]; exists {
						dist := distance(0, 0, x+1, y)
						if dist < minDistance {
							minDistance = dist
						}
					}
				}
				x += 1
			}
		}
	}
	return minDistance
}

func main() {
	wire1Input, wire2Input := readWireInputs("input.txt")
	coordinateMap := make(map[int]map[int]struct{}, len(wire1Input))
	populateCoordinates(coordinateMap, wire1Input)
	minDistance := getSmallestDistance(coordinateMap, wire2Input)
	fmt.Printf("Part 1: Minimum Distance: %d\n", minDistance)


}