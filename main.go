package main

import (
	"adventOfCode2022/day1"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Enter Day: ")
	reader := bufio.NewReader(os.Stdin)
	dayInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't read day input")
		return
	}

	day := strings.TrimSuffix(dayInput, "\n")
	dayNumber, err := strconv.Atoi(day)
	if err != nil {
		panic(err)
	}

	puzzleInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't read puzzle input")
		return
	}

	puzzle := strings.TrimSuffix(puzzleInput, "\n")
	puzzleNumber, err := strconv.Atoi(puzzle)
	if err != nil {
		panic(err)
	}

	findMatchDayPuzzle(dayNumber, puzzleNumber)
}

func findMatchDayPuzzle(day, puzzle int) {
	switch day {
	case 1:
		switch puzzle {
		case 1:
			day1.Day1Puzzle1()
		case 2:
			day1.Day1Puzzle2()
		}

	}
}
