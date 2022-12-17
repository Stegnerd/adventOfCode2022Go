package day4

import (
	"adventOfCode2022/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Puzzle1() {
	file := util.ReadFile("day4puzzle1data.txt")

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		ranges := strings.Split(row, ",")

		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		if containCheck(range1, range2) {
			count += 1
		}
	}

	fmt.Printf("\n\n count: %d \n\n", count)
}

func containCheck(range1, range2 []string) bool {
	range1Start, _ := strconv.Atoi(range1[0])
	range1End, _ := strconv.Atoi(range1[1])
	range2Start, _ := strconv.Atoi(range2[0])
	range2End, _ := strconv.Atoi(range2[1])

	if range1Start <= range2Start && range1End >= range2End {
		return true
	}

	if range2Start <= range1Start && range2End >= range1End {
		return true
	}

	return false
}

func Puzzle2() {
	file, err := os.Open("data/day4puzzle1data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		ranges := strings.Split(row, ",")

		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		if containCheck(range1, range2) {
			count += 1
		}
	}

}
