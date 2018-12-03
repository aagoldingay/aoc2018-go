package main

import (
	"bufio"
	"fmt"
	"os"
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
	//claims := []claim{}

	s.Scan()
	fmt.Printf("output : %v\n", strings.Fields(s.Text()))

	// for s.Scan() {
	// 	row := s.Text()

	// 	// boxes = append(boxes, row)
	// }

	// return claims
	return []claim{}
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
	//claims := testInput()
	fmt.Println(claims)
	//fmt.Println(part1(claims))
}

func part1(claims []claim) int {
	fmt.Printf("claims: %v\n", claims)
	overlaps := make(map[coord]int)
	fabric := [][]int{}

	for i := 0; i < 8; i++ { // change to 1000 for real thing
		fabric = append(fabric, make([]int, 8)) //change to 1000 for real thing
	}
	fmt.Printf("len : %v\n", len(fabric))
	fmt.Printf("width : %v\n", len(fabric[0]))

	for _, c := range claims {
		fmt.Printf("claim : %v\n", c)
		for i := c.fromTop; i < c.height+1; i++ {
			for j := c.fromLeft; j < c.width+1; j++ {
				fmt.Printf("coord : xy(%v , %v)\n", i, j)
				if fabric[i][j] > 0 {
					overlaps[coord{j, i}]++
				}
				fabric[i][j] = c.claimID
			}
		}
	}
	fmt.Println(overlaps)
	return len(overlaps)
	// some data structure

	// to draw
	// #123 @3,2: 5x4
	// 3 in from left 	, 	2 down from top
	// [j]				,	[i]
	// i := 2; i < 4; i++
	// 		j := 3; j < 5; j++
	//		add claim number
	//

	// if hasNumber { 'X' }
	// 		var++
	//output var
}
