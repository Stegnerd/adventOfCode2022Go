package day4

import (
	"adventOfCode2022/util"
	"bufio"
	"fmt"
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
	file := util.ReadFile("day4puzzle1data.txt")

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		ranges := strings.Split(row, ",")

		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		if containCheckPartial(range1, range2) {
			count += 1
		}
	}

	fmt.Printf("\n\n count: %d \n\n", count)
}

func containCheckPartial(range1, range2 []string) bool {
	range1Start, _ := strconv.Atoi(range1[0])
	range1End, _ := strconv.Atoi(range1[1])
	range2Start, _ := strconv.Atoi(range2[0])
	range2End, _ := strconv.Atoi(range2[1])

	// total contains
	if range1Start <= range2Start && range1End >= range2End {
		return true
	}

	// total contains
	if range2Start <= range1Start && range2End >= range1End {
		return true
	}

	// partial
	//	2-6,4-8 | 2 <= 4 | 2 <= 8
	// r2 overlaps partial of r1
	if range2Start <= range1Start && range1Start <= range2End {
		return true
	}

	// 5-7,7-9 | 7 <= 7 | 7 <= 9
	// r1 overlaps r2 in single section
	if range2Start <= range1End && range1End <= range2End {
		return true
	}

	// 2-8,3-7 | 2 <= 3 | 3 <= 8
	// r1 overlaps partial of r2.
	if range1Start <= range2Start && range2Start <= range1End {
		return true
	}

	//  6-6,4-6|  6 <= 6 | 6 <= 6
	// r1 overlaps r2 in a single section
	if range1Start <= range2End && range2End <= range1End {
		return true
	}

	return false
}
