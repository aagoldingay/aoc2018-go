package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput() []int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	nums := []int{}

	for s.Scan() {
		row, _ := strconv.Atoi(s.Text())
		nums = append(nums, row)
	}

	return nums
}

func main() {
	nums := readInput()

	part1(nums)
	fmt.Printf("p2 output : %v", part2(nums))
}

func part1(nums []int) {
	freq := 0

	for i := 0; i < len(nums); i++ {
		freq += nums[i]
	}
	fmt.Printf("p1 output : %v\n", freq)
}

func part2(nums []int) int {
	freq := 0
	values := make(map[int]bool)

	for { // needs to run over the list multiple times
		for i := 0; i < len(nums); i++ {
			freq += nums[i]
			_, prs := values[freq]
			if prs {
				return freq
			}
			values[freq] = true
		}
	}
	return -43
}
