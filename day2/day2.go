package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	A int = 1 // rock
	B     = 2 // paper
	C     = 3 // scissors
	X int = 1 // rock
	Y     = 2 // paper
	Z     = 3 // scissors
)

var (
	opponentGuesses = map[string]int{
		"A": A, // rock
		"B": B, // paper
		"C": C, // scissors
	}

	myGuesses = map[string]int{
		"X": X, // rock
		"Y": Y, // paper
		"Z": Z, // scissors
	}
)

func Puzzle1() {
	file, err := os.Open("data/day2puzzle1data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		rowArr := strings.Split(row, " ")

		opponentGuess := opponentGuesses[rowArr[0]]
		myGuess := myGuesses[rowArr[1]]

		result += matchResult(opponentGuess, myGuess) + myGuess

		fmt.Printf("\n here is your score: %d \n", result)
	}
}

func matchResult(opponent, me int) int {
	score := 0
	switch opponent {
	case A:
		if me == X {
			score = 3
		} else if me == Y {
			score = 6
		} else {
			score = 0
		}
	case B:
		if me == X {
			score = 0
		} else if me == Y {
			score = 3
		} else {
			score = 6
		}
	case C:
		if me == X {
			score = 6
		} else if me == Y {
			score = 0
		} else {
			score = 3
		}
	}

	return score
}

func Puzzle2() {

}
