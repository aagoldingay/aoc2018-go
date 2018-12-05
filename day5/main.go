package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func readInput() []byte { // could do just a string?
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return input
}

func main() {
	input := readInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(poly []byte) int {
	// poly = []byte("abBA")
	// poly = []byte("dabAcCaCBAcCcaDA")
	// fmt.Println(string(poly))
	i := 0
	for {
		p1, p2 := string(poly[i]), string(poly[i+1])

		if (strings.ToLower(p1) == p2 && strings.ToUpper(p2) == p1) || strings.ToLower(p2) == p1 && strings.ToUpper(p1) == p2 {
			poly = append(poly[:i], poly[i+2:]...)
			i--
			if i < 0 {
				i = 0
			}
			// fmt.Println(string(poly))
		} else {
			i++
		}
		if i >= len(poly)-1 {
			break
		}
	}
	return len(poly)
}

func part2(poly []byte) int {
	// poly = []byte("dabAcCaCBAcCcaDA")
	bestLength := math.MaxInt64
	for i := 0; i < 26; i++ { //letters of the alphabet
		cU, cL := i+65, i+97 // 'A', 'a'
		p := strings.Replace(string(poly), string(cU), "", -1)
		p = strings.Replace(string(p), string(cL), "", -1)
		cLength := part1([]byte(p))
		fmt.Printf("%v:%v\n", string(cU), cLength)
		if cLength < bestLength {
			bestLength = cLength
		}
	}
	return bestLength // incorrect, but test string passes
}
