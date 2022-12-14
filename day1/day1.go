package day1

import (
	"adventOfCode2022/util"
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
)

func Puzzle1() {
	file := util.ReadFile("day1puzzle1data.txt")

	highScore := 0
	currentScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			if currentScore > highScore {
				highScore = currentScore
			}
			currentScore = 0
		} else {
			rowNumber, err := strconv.Atoi(row)
			if err != nil {
				log.Fatal("Failed to convert row", row)
			}
			currentScore += rowNumber
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("High Score is: %d", highScore)
}

func Puzzle2() {
	file := util.ReadFile("day1puzzle1data.txt")

	scoreList := []int{}
	currentScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			scoreList = append(scoreList, currentScore)
			currentScore = 0
		} else {
			rowNumber, err := strconv.Atoi(row)
			if err != nil {
				log.Fatal("Failed to convert row", row)
			}
			currentScore += rowNumber
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(scoreList, func(x, y int) bool {
		return scoreList[x] > scoreList[y]
	})

	totalTop3 := scoreList[0] + scoreList[1] + scoreList[2]

	fmt.Printf("Sum of top 3: %d", totalTop3)
}
