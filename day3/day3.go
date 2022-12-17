package day3

import (
	"adventOfCode2022/util"
	"bufio"
	"fmt"
)

var (
	charValue = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
)

func Puzzle1() {
	file := util.ReadFile("day3puzzle1data.txt")

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		lenOfString := len(row) / 2
		firstHalf := row[0:lenOfString]
		secondHalf := row[lenOfString:]

		result += charCheck(firstHalf, secondHalf, lenOfString)
	}

	fmt.Printf("\n result:= %d \n", result)
}

func charCheck(firstHalf, secondHalf string, count int) int {

	wordOneCharCount := map[string]int{}
	wordTwoCharCount := map[string]int{}

	for i := 0; i < count; i++ {
		oneChar := string(firstHalf[i])
		twoChar := string(secondHalf[i])

		wordOneCharCount[oneChar] += 1
		wordTwoCharCount[twoChar] += 1
	}

	result := 0
	for k, v := range wordTwoCharCount {

		oneValue := v
		secondVal := wordOneCharCount[k]

		if oneValue > 0 && secondVal > 0 {
			result = charValue[k]
		}
	}

	return result
}

func Puzzle2() {
	file := util.ReadFile("day3puzzle1data.txt")

	index := 0
	groupCount := 1
	group := []string{}
	groupMap := map[int][]string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		group = append(group, row)

		if groupCount%3 == 0 {
			groupMap[index] = group
			group = []string{}
			groupCount = 1
			index += 1
		} else {
			groupCount += 1
		}
	}

	score := groupCheck(groupMap)

	fmt.Printf("\n result:= %d \n", score)
}

func groupCheck(groupMap map[int][]string) int {
	result := 0

	for _, group := range groupMap {
		oneMap := getCharMap(group[0])
		twoMap := getCharMap(group[1])
		threeMap := getCharMap(group[2])

		result += badgeCompare(oneMap, twoMap, threeMap)
	}

	return result
}

func getCharMap(group string) map[string]int {
	charMap := map[string]int{}

	for i := 0; i < len(group); i++ {
		oneChar := string(group[i])

		charMap[oneChar] += 1
	}

	return charMap
}

func badgeCompare(oneMap, twoMap, threeMap map[string]int) int {
	result := 0
	for k, v := range oneMap {

		oneValue := v
		secondVal := twoMap[k]
		thirdVal := threeMap[k]

		if oneValue > 0 && secondVal > 0 && thirdVal > 0 {
			result = charValue[k]
		}
	}

	return result
}
