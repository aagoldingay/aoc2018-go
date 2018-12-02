package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	boxes := []string{}

	for s.Scan() {
		row := s.Text()
		boxes = append(boxes, row)
	}

	return boxes
}

func main() {
	data := readInput()
	fmt.Printf("result: %v\n", part1(data))
	s1, s2 := part2(data)
	fmt.Printf("result: %v,%v\n", s1, s2)
}

func part1(data []string) int {
	var twoLetters, threeLetters int

	for i := 0; i < len(data); i++ {
		chars := make(map[byte]int)
		for j := 0; j < len(data[i]); j++ {
			_, prs := chars[data[i][j]]
			if !prs {
				chars[data[i][j]] = 1
				continue
			}
			chars[data[i][j]]++
		}

		twoFound, threeFound := false, false
		for _, v := range chars {
			if v == 2 && !twoFound {
				twoLetters++
				twoFound = true
			}
			if v == 3 && !threeFound {
				threeLetters++
				threeFound = true
			}
		}
	}
	return twoLetters * threeLetters
}

func part2(data []string) (string, string) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if stringCheck(data[i], data[j]) {
				return data[i], data[j]
			}
		}
	}
	return "none", "none"
}

func stringCheck(one, two string) bool {
	mistakesAllowed := 1
	if len(one) == len(two) {
		for i := 0; i < len(one); i++ {
			if one[i] != two[i] {
				mistakesAllowed--
				if mistakesAllowed < 0 {
					return false
				}
			}
		}
	}
	return true
}
