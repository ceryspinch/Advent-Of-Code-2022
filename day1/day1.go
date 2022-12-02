package main

import (
	fileReader "advent-of-code/utils"
	"fmt"
	"strconv"
)

func calculateTotalCalories() ([]int, error) {
	inputs, err := fileReader.ReadInput()
	if err != nil {
		fmt.Println("There was an error reading the input file.")
		return nil, err
	}

	elves := make([]int, 10000)
	i := 0
	num := 0
	totalCalories := 0

	// Range over each line of the input file
	for _, line := range inputs {
		// If the line is not empty (so still on the same elf), convert to an int
		// and add to that elf's total
		if line != "" {
			num, err = strconv.Atoi(line)
			if err != nil {
				fmt.Println("There was an error converting the string to an integer.")
				return nil, err
			}
			totalCalories = totalCalories + num
		} else {
			// When finished with an elf, add their total calories to an array
			elves[i] = totalCalories

			// Reset totalCalories ready for the next elf
			totalCalories = 0
			i++
		}
	}
	return elves, nil
}

func PuzzleOne() (int, error) {
	totals, err := calculateTotalCalories()
	if err != nil {
		fmt.Println("There was an error getting the total calories for each elf.")
		return 0, err
	}

	// Find largest number in array
	max := 0
	for _, num := range totals {
		if num > max {
			max = num
		}
	}

	return max, nil
}

func PuzzleTwo() (int, error) {

	totals, err := calculateTotalCalories()
	if err != nil {
		fmt.Println("There was an error getting the total calories for each elf.")
		return 0, err
	}
	first := 0
	second := 0
	third := 0

	// Find elf with largest total calories
	for _, calories := range totals {
		if calories > first {
			first = calories
		}
	}

	// Find elf with second largest total calories
	for _, calories := range totals {
		if calories > second && calories < first {
			second = calories
		}
	}

	// Find elf with third largest total calories
	for _, calories := range totals {
		if calories > third && calories < second {
			third = calories
		}
	}

	// Return total calories of top 3 elves
	return (first + second + third), nil
}

// func main() {
// 	solutionOne, err := PuzzleOne()
// 	if err != nil {
// 		fmt.Println("There was an error getting the solution to puzzle one.")
// 	}


// 	solutionTwo, err := PuzzleTwo()
// 	if err != nil {
// 		fmt.Println("There was an error getting the solution to puzzle two.")
// 	}

// 	// Print solutions
// 	fmt.Println("The solution to puzzle one is: ", solutionOne)
// 	fmt.Println("The solution to puzzle two is: ", solutionTwo)
// }



