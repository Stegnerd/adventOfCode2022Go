package main

import (
	"adventOfCode2022/day1"
	"adventOfCode2022/day2"
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
		switchPuzzles(puzzle, day1.Puzzle1, day1.Puzzle2)
	case 2:
		switchPuzzles(puzzle, day2.Puzzle1, day2.Puzzle2)
	}
}

func switchPuzzles(choice int, puzzle1 func(), puzzle2 func()) {
	switch choice {
	case 1:
		puzzle1()
	case 2:
		puzzle2()
	}
}
