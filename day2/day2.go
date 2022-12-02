package main

import (
	fileReader "advent-of-code/utils"
	"fmt"
	"strings"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3

	loss = 0
	draw = 3
	win  = 6
)

func PuzzleOne() (int, error) {
	inputList, err := fileReader.ReadInput()

	if err != nil {
		fmt.Println("There was an error reading the input file.")
		return 0, err
	}

	totalScore := 0

	for _, line := range inputList {
		sections := strings.Split(line, " ")
		opponent := sections[0]
		me := sections[1]

		if opponent == "A" {
			if me == "X" {
				totalScore += (rock + draw)
			} else if me == "Y" {
				totalScore += (paper + win)
			} else {
				totalScore += (scissors + loss)
			}
		} else if opponent == "B" {
			if me == "X" {
				totalScore += (rock + loss)
			} else if me == "Y" {
				totalScore += (paper + draw)
			} else {
				totalScore += (scissors + win)
			}
		} else {
			if me == "X" {
				totalScore += (rock + win)
			} else if me == "Y" {
				totalScore += (paper + loss)
			} else {
				totalScore += (scissors + draw)
			}
		}
	}

	return totalScore, nil
}

func PuzzleTwo() (int, error) {
	inputList, err := fileReader.ReadInput()

	if err != nil {
		fmt.Println("There was an error reading the input file.")
		return 0, err
	}

	totalScore := 0

	for _, line := range inputList {
		sections := strings.Split(line, " ")
		opponent := sections[0]
		me := sections[1]

		if opponent == "A" {
			if me == "X" {
				totalScore += (scissors + loss)
			} else if me == "Y" {
				totalScore += (rock + draw)
			} else {
				totalScore += (paper + win)
			}
		} else if opponent == "B" {
			if me == "X" {
				totalScore += (rock + loss)
			} else if me == "Y" {
				totalScore += (paper + draw)
			} else {
				totalScore += (scissors + win)
			}
		} else {
			if me == "X" {
				totalScore += (paper + loss)
			} else if me == "Y" {
				totalScore += (scissors + draw)
			} else {
				totalScore += (rock + win)
			}
		}
	}

	return totalScore, nil
}

func main() {
	sol1, err := PuzzleOne()
	if err != nil {
		fmt.Println("There was an error getting the solution to puzzle one.")
	}

	fmt.Println(sol1)

	sol2, err := PuzzleTwo()
	if err != nil {
		fmt.Println("There was an error getting the solution to puzzle two.")
	}

	fmt.Println(sol2)

}
