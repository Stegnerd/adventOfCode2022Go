package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	file, err := os.Open("data/day3puzzle1data.txt")
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

}
