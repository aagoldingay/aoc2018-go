package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	claimID, fromLeft, fromTop, width, height int
}

type coord struct {
	x, y int
}

func readInput() []claim {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	claims := []claim{}

	for s.Scan() {
		info := strings.Fields(s.Text())
		id, _ := strconv.Atoi(info[0][1:])

		coords := strings.Split(info[2][:len(info[2])-1], ",")
		dimensions := strings.Split(info[3], "x")

		coordLeft, _ := strconv.Atoi(coords[0])
		coordTop, _ := strconv.Atoi(coords[1])
		width, _ := strconv.Atoi(dimensions[0])
		height, _ := strconv.Atoi(dimensions[1])
		claims = append(claims, claim{id, coordLeft, coordTop, width, height})
	}

	return claims
}

func testInput() []claim {
	// #1 @ 1,3: 4x4
	// #2 @ 3,1: 4x4
	// #3 @ 5,5: 2x2

	return []claim{
		{1, 1, 3, 4, 4},
		{2, 3, 1, 4, 4},
		{3, 5, 5, 2, 2},
	}
}

func main() {
	claims := readInput()
	// claims := testInput()
	fmt.Println(parts(claims))
}

func parts(claims []claim) int {
	//part 1
	overlaps := make(map[coord]int)
	fabric := make(map[coord]int)

	for _, c := range claims {
		for i := 0; i < c.height; i++ {
			for j := 0; j < c.width; j++ {
				cd := coord{i + c.fromTop, j + c.fromLeft}
				if _, prs := fabric[cd]; prs {
					overlaps[cd]++
					fabric[cd] = -1
					continue
				}
				fabric[cd] = c.claimID
			}
		}
	}

	//part 2
	for _, c := range claims {
		area := c.height * c.width
		var currentArea int

		for _, id := range fabric {
			if id == c.claimID {
				currentArea++
			}
		}

		if currentArea == area {
			fmt.Printf("claim id: %v\n", c.claimID)
		}
	}

	//returns part 1 answer
	return len(overlaps)
}
